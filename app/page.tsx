"use client";
import Image from "next/image";
import { useState } from "react";

export default function Home() {
  const [quote, setQuote] = useState<string>("");
  const [copyMessage, setCopyMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  async function fetchQuote() {
    const apiUrl = "/api/quote"; // Direct API route - works in both development and production
    
    setIsLoading(true);
    try {
      const res = await fetch(apiUrl);
      const data = await res.json();
      setQuote(data.quote);
    } catch (error) {
      setQuote("Failed to fetch quote. Please try again!");
    } finally {
      setIsLoading(false);
    }
  }

  const copyToClipboard = () => {
    navigator.clipboard.writeText(quote);
    setCopyMessage('Quote copied to clipboard!');
    setTimeout(() => setCopyMessage(''), 2000); // Clear message after 2 seconds
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-pink-400 via-purple-400 to-blue-400 px-4 py-10">
      <main className="flex flex-col gap-10 items-center w-full max-w-2xl">
        <div className="flex flex-col items-center gap-2">
          <Image
            src="/genz-logo.png"
            alt="GenZ Quote Logo"
            width={120}
            height={120}
            className="mb-2 drop-shadow-xl animate-bounce rounded-full border-4 border-white/40 hover:scale-110 transition-transform duration-300 bg-white/70"
            priority
          />
          <h1 className="text-5xl font-extrabold text-white drop-shadow-lg tracking-tight text-center font-[Poppins,Inter,sans-serif] animate-bounce">
            GenZ Quote Generator
          </h1>
        </div>
        <div className="flex flex-col items-center gap-6 w-full">
          <button
            className="px-8 py-4 bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-500 text-white font-bold text-2xl rounded-full shadow-lg hover:scale-105 hover:from-pink-500 hover:to-blue-500 transition-all duration-300 active:scale-95 border-4 border-white/30 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:scale-100"
            onClick={fetchQuote}
            disabled={isLoading}
          >
            <span className="inline-flex items-center gap-2">
              <span role="img" aria-label="sparkles">
                {isLoading ? "⏳" : "⚡"}
              </span> 
              {isLoading ? "Generating..." : "Generate Quote"}
            </span>
          </button>
          <style>{`
            @keyframes rollout {
              0% { opacity: 0; transform: translateY(40px) scaleX(0.7); }
              60% { opacity: 1; transform: translateY(-8px) scaleX(1.05); }
              100% { opacity: 1; transform: translateY(0) scaleX(1); }
            }
            .rollout {
              animation: rollout 0.7s cubic-bezier(.68,-0.55,.27,1.55);
            }
          `}</style>
          <div
            className={`text-3xl font-semibold text-white text-center min-h-[64px] max-w-xl p-6 rounded-2xl bg-white/10 backdrop-blur-lg shadow-xl border-2 border-white/20 ${quote ? 'rollout' : ''}`}
          >
            {isLoading ? (
              <div className="flex items-center justify-center gap-2">
                <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-white"></div>
                Getting your quote...
              </div>
            ) : quote ? (
              quote
            ) : (
              "Tap the button for instant GenZ inspo!"
            )}
          </div>
          {quote && (
            <button
              onClick={copyToClipboard}
              className="px-4 py-2 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-600 transition-all duration-300"
            >
              Copy Quote
            </button>
          )}
          {copyMessage && <div className="mt-2 text-green-200">{copyMessage}</div>}
        </div>
        {/* ...existing code... (keep the ol and links for demo) */}
    
       
         
            
          
        
      </main>
 
    </div>
  );
}
