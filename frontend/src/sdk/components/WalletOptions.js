import * as React from "react";
import { Connector, useConnect } from "wagmi";
import { CoinbaseWalletLogo } from "./CoinbaseWalletLogo";
import MetaMaskWalletLogo from "./MetaMaskWalletLogo";

export function WalletOptions() {
  const { connectors, connect } = useConnect();
  return connectors.map((connector) => (
    <WalletOption
      key={connector.uid}
      connector={connector}
      onClick={() => connect({ connector })}
    />
  ));
}

function WalletOption({ connector, onClick }) {
  const [ready, setReady] = React.useState(false);

  React.useEffect(() => {
    (async () => {
      const provider = await connector.getProvider();
      setReady(!!provider);
    })();
  }, [connector]);

  return (
    <div className="flex flex-col items-center justify-center my-2">
      {connector.name == "Coinbase Wallet" && (
        <div className="bg-[#0052FF] flex w-[80%]  px-2 rounded-t-lg">
          <span className="text-[10px] font-semibold">Gas less transaction!</span>
        </div>
      )}
      <button
        className={`bg-[#383838] hover:bg-[rgb(50,50,50)] active:scale-90 w-[80%] ${connector.name == "Coinbase Wallet" ? "rounded-b-lg" : "rounded-lg"}  flex p-3 items-center justify-center`}
        disabled={!ready}
        onClick={onClick}
      >
        {connector.name == "Coinbase Wallet" && <CoinbaseWalletLogo />}
        {connector.name == "MetaMask" && <MetaMaskWalletLogo />}
        <span className="text-xl font-semibold ml-3">{connector.name}</span>
      </button>
    </div>
  );
}
