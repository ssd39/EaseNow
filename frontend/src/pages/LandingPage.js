import React, { useEffect } from "react";
import { BackgroundBeams } from "./background-beams";
import { useNavigate } from "react-router-dom";
import Typewriter from "typewriter-effect/dist/core";

export default function LandingPage() {
  const navigate = useNavigate();

  useEffect(() => {
    new Typewriter("#lol-text", {
      strings: [
        "Instant Credit Using Real-World Assets Aka Your Credit Score.",
        "Autonomous Vault Security For Your Private Data Taken As Security Deposit.",
        "Smart Wallet Integration So Even You Have 0 ETH You Can Purchase Things.",
        "Buy Anytime And Settle Payments From Your Cold Wallet At Your Convenience."
      ],
      autoStart: true,
      loop: true
    });
  }, []);
  return (
    <div className="min-h-screen bg-black p-12 flex flex-col">
      <div className="flex flex-col select-none ">
        <div>
          <span className="title text-6xl font-bold">EaseNow⚡</span>
        </div>
        <span className="text-lg text-orange-300">Buy Now Pay Later</span>
      </div>
      <div className="mt-10 flex h-[75vh]">
        <div className="flex-1 flex justify-center flex-col">
          <div
            onClick={() => navigate("/manage-account")}
            className="h-[150px] w-[560px] g1bg rounded-lg p-4 flex flex-col justify-center cursor-pointer active:scale-90 select-none"
          >
            <span className="text-white text-3xl font-semibold">
              Manage Your Credits {">"}
            </span>
            <span className="text-orange-200">
              A dashboard to mange your credit usage and repay the debt
            </span>
          </div>
          <div
            onClick={() => navigate("/test-shop")}
            className="h-[150px] w-[560px] g2bg rounded-lg p-4 flex flex-col mt-4 justify-center cursor-pointer active:scale-90 select-none"
          >
            <span className="text-white text-3xl font-semibold">
              Demo Shop {">"}
            </span>
            <span className="text-orange-200">
              A demo shop to witness magical experience by paying using
              EaseNow⚡
            </span>
          </div>
          <div   onClick={() => alert("Under development. Comming soon!")} className="h-[150px] w-[560px] g3bg rounded-lg p-4 flex flex-col mt-4 justify-center cursor-pointer active:scale-90 select-none">
            <span className="text-white text-3xl font-semibold">
              Earn Extra By Lending {">"}
            </span>
            <span className="text-orange-200">
              EaseNow⚡ is backed by the smart DeFI protocol, Be part of it and
              earn ENT tokens.
            </span>
          </div>
          <div onClick={() => alert("Under development. Comming soon!")} className="h-[150px] w-[560px] g4bg rounded-lg p-4 flex flex-col mt-4 justify-center cursor-pointer active:scale-90 select-none">
            <span className="text-white text-3xl font-semibold">
              Check Public Credit Records {">"}
            </span>
            <span className="text-orange-200">
              The private data of users unable to pay their debts is decrypted
              by an autonomous data vault powered by TEE, making the previously
              encrypted information accessible.
            </span>
          </div>
        </div>
        <div className="flex-1 relative overflow-hidden border border-solid border-r-0 rounded-l-3xl border-[rgba(255,255,255,0.1)]">
          <div className="stars">
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
            <div class="star"></div>
          </div>
          <div className="w-ful h-full flex items-center justify-center p-4">
            <span
              className="text-white font-semibold text-3xl text-center"
              id="lol-text"
            >
              
            </span>
          </div>
          <div className="absolute img1 top-[10%] left-[42%]">
            <img src="/1.png" height={128} width={128} />
          </div>
          <div className="absolute img2 top-[15%] left-[10%]">
            <img src="/2.png" height={128} width={128} />
          </div>
          <div className="absolute img3 bottom-[15%] left-[8%]">
            <img src="/3.png" height={128} width={128} />
          </div>
          <div className="absolute img4 top-[15%] right-[10%]">
            <img src="/4.png" height={128} width={128} />
          </div>
          <div className="absolute img5 bottom-[15%] right-[10%]">
            <img src="/5.png" height={128} width={128} />
          </div>
          <div className="absolute img6 bottom-[10%] right-[42%]">
            <img src="/6.png" height={128} width={128} />
          </div>
        </div>
      </div>
    </div>
  );
}
