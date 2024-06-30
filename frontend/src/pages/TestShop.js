import React, { useState } from "react";
import EaseNow from "../sdk/EaseNow";

export default function TestShop() {
  const [isPurchase, setIsPurchase] = useState(false);
  return (
    <div className="w-full">
      <div className=" p-5 ">
        <span className="text-3xl font-bold">Test Shop</span>
      </div>
      <hr />
      <div className="flex mx-[200px] my-10">
        <div className="p-2 border border-solid  rounded-lg">
          <img height={512} width={512} src="/sticker.png" />
        </div>
        <div className="flex flex-col ml-8">
          <span className="text-2xl font-semibold">A Head Shaver</span>
          <span className="text-xl text-orange-700 font-semibold">
            Price: <span className="text-black font-normal">0.001 ETH</span>
          </span>
          <div className="mt-8">
            <button className="bg-[gray] text-white p-2 px-6  rounded-full active:scale-90">Buy</button>
            <button onClick={()=> setIsPurchase(true)} className="bg-[gray] text-white p-2 px-6  rounded-full ml-4 active:scale-90">Buy with EaseNow</button>
          </div>
        </div>
      </div>

      {isPurchase && (
        <EaseNow
          isOpen={isPurchase}
          handleClose={() => {
            setIsPurchase(false)
          }}
          merchent="0x2e570B3A13d6954a1b9dfeDf68981f1b79F2d090"
          amount={"0.001"}
          title="A Head Shaver"
        />
      )}
    </div>
  );
}
