import React, { useEffect, useState } from "react";
import Dialog from "@mui/material/Dialog";
import { WagmiProvider, useAccount } from "wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider, useDispatch } from "react-redux";
import { config } from "./config";
import { WalletOptions } from "./components/WalletOptions";
import Header from "./components/Header";
import store from "./redux/store";
import "./EaseNow.css";
import Payment from "./components/Payment";
import Registration from "./components/Registration";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const queryClient = new QueryClient();

export function ConnectWallet({ nextStage }) {
  const { isConnected } = useAccount();
  useEffect(() => {
    if (isConnected) {
      nextStage();
    }
  }, [isConnected]);
  return <WalletOptions />;
}

export default function EaseNow({ isOpen,  handleClose, merchent, amount, title }) {
  const [stage, setStage] = useState(0);

  return (
    <WagmiProvider config={config}>
      <ToastContainer />
      <QueryClientProvider client={queryClient}>
        <Provider store={store}>
          <Dialog
            sx={{
              ".MuiPaper-root": {
                background: "black",
                color: "white",
                borderRadius: 3,
                minHeight: 500,
                maxWidth: 400,
                width: "100%",
              },
            }}
            onClose={handleClose}
            open={isOpen}
          >
            <div className="p-3 flex flex-col flex-1">
              <Header onClose={() => handleClose(false)} />
              {false && (
                <div className="mt-6">
                  <span className="font-bold">
                    #OrderId: <span className="font-normal">1234</span>
                  </span>
                </div>
              )}
              <div className="flex flex-1">
                {stage == 0 && (
                  <div className="flex flex-col w-full">
                    <div className="items-center justify-center flex mt-6">
                      <img src="wallet.png" />
                      <span className="text-2xl ml-1">Connect a wallet</span>
                    </div>
                    <div className="mt-6">
                      <ConnectWallet nextStage={() => setStage(1)} />
                    </div>
                  </div>
                )}
                {stage == 1 && (
                  <Payment
                    onClose={() => handleClose(false)}
                    merchent={merchent}
                    amount={amount}
                    title={title}
                    nextStage={(id) => setStage(id)}
                  />
                )}
                {stage == 2 && (
                  <Registration nextStage={(id) => setStage(id)} />
                )}
              </div>
            </div>
          </Dialog>
        </Provider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
