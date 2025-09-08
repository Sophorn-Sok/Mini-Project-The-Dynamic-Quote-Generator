import { NextRequest, NextResponse } from 'next/server';

const SUPABASE_URL = process.env.SUPABASE_URL;
const SUPABASE_API_KEY = process.env.SUPABASE_API_KEY;

interface Quote {
  id: number;
  text: string;
  author: string;
}

// Fallback quotes if Supabase is not available
const fallbackQuotes: Quote[] = [
  { id: 1, text: "✨ The best way to get started is to quit talking and begin doing.", author: "Walt Disney" },
  { id: 2, text: "🔥 Don't let yesterday take up too much of today.", author: "Will Rogers" },
  { id: 3, text: "💪 It's not whether you get knocked down, it's whether you get up.", author: "Vince Lombardi" },
  { id: 4, text: "🚀 If you are working on something exciting, it will keep you motivated.", author: "GenZ Wisdom" },
  { id: 5, text: "🌈 Success is not in what you have, but who you are.", author: "Bo Bennett" },
  { id: 6, text: "😎 Dream big, hustle harder.", author: "GenZ Motivation" },
  { id: 7, text: "👾 Stay weird, stay creative.", author: "GenZ Vibes" },
  { id: 8, text: "🦄 Be yourself, everyone else is taken.", author: "Oscar Wilde" },
  { id: 9, text: "💥 Make it happen, Gen Z style!", author: "GenZ Energy" },
  { id: 10, text: "🌟 You are the main character of your story.", author: "GenZ Wisdom" },
];

// Cache for quotes - will store quotes in memory
let quotesCache: Quote[] | null = null;
let cacheTimestamp = 0;
const CACHE_DURATION = 5 * 60 * 1000; // 5 minutes in milliseconds

async function getQuotesFromSupabase(): Promise<Quote[] | null> {
  // Check if we have valid cached data
  const now = Date.now();
  if (quotesCache && (now - cacheTimestamp) < CACHE_DURATION) {
    console.log(`Using cached quotes (${quotesCache.length} quotes)`);
    return quotesCache;
  }

  if (!SUPABASE_URL || !SUPABASE_API_KEY) {
    return null;
  }

  try {
    console.log('Fetching fresh quotes from Supabase...');
    const response = await fetch(`${SUPABASE_URL}/rest/v1/quotes?select=*`, {
      headers: {
        'apikey': SUPABASE_API_KEY,
        'Authorization': `Bearer ${SUPABASE_API_KEY}`,
      },
    });

    if (!response.ok) {
      return null;
    }

    const quotes: Quote[] = await response.json();
    if (quotes.length > 0) {
      // Update cache
      quotesCache = quotes;
      cacheTimestamp = now;
      console.log(`Cached ${quotes.length} quotes from Supabase`);
      return quotes;
    }
    return null;
  } catch (error) {
    console.error('Supabase error:', error);
    return null;
  }
}

export async function GET() {
  try {
    // Try to get quotes from cache or Supabase
    let quotes = await getQuotesFromSupabase();
    
    if (!quotes) {
      console.log('Using fallback quotes');
      quotes = fallbackQuotes;
    }

    // Pick a random quote - this is now instant!
    const randomIndex = Math.floor(Math.random() * quotes.length);
    const selectedQuote = quotes[randomIndex];

    return NextResponse.json({
      quote: selectedQuote.text,
      author: selectedQuote.author,
    });
  } catch (error) {
    console.error('Error in quote API:', error);
    
    // Return a fallback quote even if there's an error
    const randomIndex = Math.floor(Math.random() * fallbackQuotes.length);
    const selectedQuote = fallbackQuotes[randomIndex];
    
    return NextResponse.json({
      quote: selectedQuote.text,
      author: selectedQuote.author,
    });
  }
}

export async function POST(request: NextRequest) {
  if (!SUPABASE_URL || !SUPABASE_API_KEY) {
    return NextResponse.json({ error: 'Supabase not configured' }, { status: 500 });
  }

  try {
    const body = await request.json();
    const { text, author } = body;

    const response = await fetch(`${SUPABASE_URL}/rest/v1/quotes`, {
      method: 'POST',
      headers: {
        'apikey': SUPABASE_API_KEY,
        'Authorization': `Bearer ${SUPABASE_API_KEY}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ text, author }),
    });

    if (!response.ok) {
      throw new Error('Failed to create quote in Supabase');
    }

    const newQuote = await response.json();
    return NextResponse.json(newQuote);
  } catch (error) {
    console.error('Error creating quote:', error);
    return NextResponse.json({ error: 'Failed to create quote' }, { status: 500 });
  }
}
