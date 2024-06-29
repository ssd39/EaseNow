import { http, createConfig } from 'wagmi';
import { baseSepolia } from 'wagmi/chains';
import { coinbaseWallet } from 'wagmi/connectors';
import { metaMask } from 'wagmi/connectors';

export const config = createConfig({
  chains: [baseSepolia],
  connectors: [
    coinbaseWallet({
      appName: 'EaseNow⚡',
      preference: 'smartWalletOnly',
    }),
    metaMask()
  ],
  transports: {
    [baseSepolia.id]: http(),
  },
});
 
