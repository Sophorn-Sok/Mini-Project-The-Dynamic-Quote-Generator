package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	supabaseUrl = os.Getenv("SUPABASE_URL")
	supabaseKey = os.Getenv("SUPABASE_API_KEY")
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// Fallback quotes if Supabase is not available
var fallbackQuotes = []Quote{
	{ID: 1, Text: "✨ The best way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
	{ID: 2, Text: "🔥 Don't let yesterday take up too much of today.", Author: "Will Rogers"},
	{ID: 3, Text: "💪 It's not whether you get knocked down, it's whether you get up.", Author: "Vince Lombardi"},
	{ID: 4, Text: "🚀 If you are working on something exciting, it will keep you motivated.", Author: "GenZ Wisdom"},
	{ID: 5, Text: "🌈 Success is not in what you have, but who you are.", Author: "Bo Bennett"},
	{ID: 6, Text: "😎 Dream big, hustle harder.", Author: "GenZ Motivation"},
	{ID: 7, Text: "👾 Stay weird, stay creative.", Author: "GenZ Vibes"},
	{ID: 8, Text: "🦄 Be yourself, everyone else is taken.", Author: "Oscar Wilde"},
	{ID: 9, Text: "💥 Make it happen, Gen Z style!", Author: "GenZ Energy"},
	{ID: 10, Text: "🌟 You are the main character of your story.", Author: "GenZ Wisdom"},
}

func getQuotesFromSupabase() ([]Quote, error) {
	if supabaseUrl == "" || supabaseKey == "" {
		return nil, nil // No Supabase config
	}

	req, err := http.NewRequest("GET", supabaseUrl+"/rest/v1/quotes?select=*", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, nil // Table might not exist yet
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func createQuoteInSupabase(quote Quote) error {
	if supabaseUrl == "" || supabaseKey == "" {
		return nil // No Supabase config
	}

	data, err := json.Marshal(quote)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", supabaseUrl+"/rest/v1/quotes", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	log.Printf("🔗 Connecting to Supabase: %s", supabaseUrl)

	// Handle /api/quote endpoint
	http.HandleFunc("/api/quote", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Try to get quotes from Supabase first
		quotes, err := getQuotesFromSupabase()
		if err != nil || len(quotes) == 0 {
			log.Printf("Using fallback quotes (Supabase error: %v)", err)
			quotes = fallbackQuotes
		} else {
			log.Printf("Retrieved %d quotes from Supabase", len(quotes))
		}

		// Pick a random quote
		localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := localRand.Intn(len(quotes))

		respObj := map[string]string{
			"quote":  quotes[randomIndex].Text,
			"author": quotes[randomIndex].Author,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respObj)
	})

	// Handle /api/quotes endpoint for CRUD operations
	http.HandleFunc("/api/quotes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		switch r.Method {
		case "GET":
			quotes, err := getQuotesFromSupabase()
			if err != nil || len(quotes) == 0 {
				// Return fallback quotes as JSON
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(fallbackQuotes)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(quotes)

		case "POST":
			var quote Quote
			err := json.NewDecoder(r.Body).Decode(&quote)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			err = createQuoteInSupabase(quote)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(quote)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("🚀 GenZ Quote API running at http://localhost:8080")
	log.Println("📌 Endpoints:")
	log.Println("   - GET /api/quote (random quote)")
	log.Println("   - GET /api/quotes (all quotes)")
	log.Println("   - POST /api/quotes (create quote)")
	http.ListenAndServe(":8080", nil)
}
