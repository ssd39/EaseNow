/* global BigInt */
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import {
  fetchUser,
  setTransactionData,
} from "../redux/actionSlice";
import { coinbaseSmartWalletProxyBytecode, contractAddr } from "../constants";
import {
  useAccount,
  usePublicClient,
  useReadContract,
  useWriteContract,
} from "wagmi";
import { CircularProgress } from "@mui/material";
import GButton from "./GButton";
import { wieToEthStr } from "../helper/network";
import abi from "../helper/abi";
import { ethers } from "ethers";
import testShopAbi from "../helper/testShopAbi";

export default function Payment({ onClose, nextStage, merchent, amount, title }) {
  const { address, connector } = useAccount();
  const isNewUser = useSelector((state) => state.action.isNewUser);
  const isUserFetched = useSelector((state) => state.action.isUserFetched);
  const [loading, setLoading] = useState(true);
  const dispatch = useDispatch();
  const { data } = useReadContract({
    address: contractAddr,
    abi,
    functionName: "getBorrowerData",
    args: [address],
  });
  const { data: hash, writeContract, isPending, isSuccess, error } = useWriteContract();

  const [btnLoadingg, setBtnLoading] = useState(false);
  const client = usePublicClient();
  useEffect(() => {
    dispatch(
      setTransactionData({
        merchent,
        amount,
        title,
      })
    );
  }, [merchent, amount, title]);

  useEffect(() => {
    if (!isUserFetched) {
      dispatch(fetchUser(address));
    }
  }, []);

  useEffect(() => {
    if (isUserFetched) {
      if (isNewUser) {
        nextStage(2);
      }
      setLoading(false);
    }
  }, [isUserFetched]);

  useEffect(() => {
    if (hash) {
      setLoading(false);
    }
  }, [hash]);

  return (
    <div className="flex flex-1">
      {loading && (
        <div className="text-blue-600 flex items-center justify-center flex-1">
          <CircularProgress size={64} />
        </div>
      )}
      {!loading && !isSuccess && (
        <>
          <div className="mt-4 flex flex-col w-full">
            <div className="flex flex-col">
              <span className="text-2xl font-semibold">Checkout 🛒</span>
              <div className="mt-4  flex justify-between items-center">
                <span className="text-xl">
                  <span className="text-[#FC466B] font-semibold">#Item:</span>{" "}
                  {title}
                </span>
                <span>{amount} ETH</span>
              </div>
            </div>
            <div className="flex-1 flex flex-col justify-end items-end my-4">
              <div className="flex">
                <div className="w-[140px]">
                  <span className="text-lg text-white">Items:</span>
                </div>
                <div className="w-[50px] flex justify-end">
                  <span className="text-lg text-[#3F5EFB] font-semibold">
                    1
                  </span>
                </div>
              </div>
              <div className="flex">
                <div className="w-[140px]">
                  <span className="text-lg text-white">Available credits:</span>
                </div>
                <div className="w-[135px] flex justify-end">
                  <span className="text-lg text-[#3F5EFB] font-semibold ml-4">
                    {wieToEthStr(data.creditLimit - data.debt)} ETH
                  </span>
                </div>
              </div>

              <div className="mt-2 w-full">
                <hr />
              </div>
              <div className="flex">
                <div className="w-[70px]">
                  <span className="text-lg text-white">Subtotal:</span>
                </div>
                <div className="w-[140px] flex justify-end">
                  <span className="text-lg text-red-400 font-semibold ml-4">
                    -{amount} ETH
                  </span>
                </div>
              </div>
            </div>
            {error && (
        <div>Error: {(error ).shortMessage || error.message}</div>
      )}
            <div className="justify-center flex mb-6">
              <GButton
                loading={isPending || btnLoadingg}
                onClick={async () => {
                  setBtnLoading(true);
                  const code = await client.getCode({ address });
                  if (!code || code !=  coinbaseSmartWalletProxyBytecode) {
                    writeContract({
                      address: contractAddr,
                      abi,
                      functionName: "borrow",
                      args: [
                        ethers.parseUnits(amount, "ether"),
                        "",
                        merchent,
                        true,
                        new ethers.Interface(testShopAbi).encodeFunctionData(
                          "buy",
                          [address]
                        ),
                      ],
                    });
                    setBtnLoading(false);
                  } else {
                    writeContract({
                      address: contractAddr,
                      abi,
                      functionName: "borrow",
                      args: [
                        ethers.parseUnits(amount, "ether"),
                        "",
                        merchent,
                        true,
                        new ethers.Interface(testShopAbi).encodeFunctionData(
                          "buy",
                          [address]
                        ),
                      ],
                    });
                    setBtnLoading(false);
                  }
                }}
              >
                Pay
              </GButton>
            </div>
          </div>
        </>
      )}
      { isSuccess  && (  <div className="mt-4 flex flex-col items-center w-full">
          <span className="text-lg text-center">
            Items purchased successfully ✅
          </span>
          <button
            onClick={() => {
              window.open(`https://base-sepolia.blockscout.com/tx/${hash}`)
            }}
            className="bg-[#383838] hover:bg-[rgb(50,50,50)]  active:scale-90  rounded-lg  w-[80%] flex p-3 items-center justify-center mt-3"
          >
            Checkout on transaction explorer
          </button>
          <div className="flex-1 items-end flex mb-6">
            <GButton
              loading={loading}
              onClick={async () => {
                onClose()
              }}
            >
              Close
            </GButton>
          </div>
        </div>)}
    </div>
  );
}
