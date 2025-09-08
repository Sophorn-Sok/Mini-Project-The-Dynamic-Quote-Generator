package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"✨ The best way to get started is to quit talking and begin doing.",
	"🔥 Don't let yesterday take up too much of today.",
	"💪 It's not whether you get knocked down, it's whether you get up.",
	"🚀 If you are working on something exciting, it will keep you motivated.",
	"🌈 Success is not in what you have, but who you are.",
	"😎 Dream big, hustle harder.",
	"👾 Stay weird, stay creative.",
	"🦄 Be yourself, everyone else is taken.",
	"💥 Make it happen, Gen Z style!",
	"🌟 You are the main character of your story.",
	"🎧 Good vibes only.",
	"💡 Think different, act bold.",
	"🫶 Spread kindness like confetti.",
	"📱 Disconnect to reconnect.",
	"🕺 Dance like nobody's watching.",
	"🍀 Luck is when preparation meets opportunity.",
	"🧠 Mindset is everything.",
	"🔥 Hustle in silence, let success make the noise.",
	"🌊 Go with the flow, but make waves.",
	"🎨 Create your own reality.",
	"💬 Speak your truth.",
	"🌻 Grow through what you go through.",
	"🦋 Change is beautiful.",
	"🎲 Take risks, regret nothing.",
	"💎 Shine bright, even on cloudy days.",
	"🌌 The future belongs to those who believe in the beauty of their dreams.",
	"🌍 Your limitation—it's only your imagination.",
	"🌠 Push yourself, because no one else is going to do it for you.",
	"🌻 Great things never come from comfort zones.",
	"🌈 Dream it. Wish it. Do it.",
	"💪 Success doesn’t just find you. You have to go out and get it.",
	"🌟 The harder you work for something, the greater you’ll feel when you achieve it.",
	"🌊 Dream bigger. Do bigger.",
	"💡 Don’t stop when you’re tired. Stop when you’re done.",
	"🔥 Wake up with determination. Go to bed with satisfaction.",
	"🌻 Do something today that your future self will thank you for.",
	"🌌 Little things make big days.",
	"🌈 It’s going to be hard, but hard does not mean impossible.",
	"💎 Push yourself, because no one else is going to do it for you.",
}

func main() {
	http.HandleFunc("/api/quote", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := localRand.Intn(len(quotes))
		resp := map[string]string{"quote": quotes[randomIndex]}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("🚀 GenZ Quote API running at http://localhost:8080/api/quote")
	http.ListenAndServe(":8080", nil)
}
