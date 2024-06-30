import React, { useEffect, useState } from "react";
import { useAccount, useReadContract } from "wagmi";
import { ConnectWallet } from "../sdk/EaseNow";
import Header from "../sdk/components/Header";
import { contractAddr } from "../sdk/constants";
import abi from "../sdk/helper/abi";
import { wieToEthStr } from "../sdk/helper/network";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { client } from "../sdk/redux/actionSlice";
import { CircularProgress } from "@mui/material";
import { gql } from "@apollo/client";

export default function Account() {
  const { isConnected, address } = useAccount();
  const { data } = useReadContract({
    address: contractAddr,
    abi,
    functionName: "getBorrowerData",
    args: [address],
  });
  const [isLoading, setLoading] = useState(false);

  const [txHistory, setHistoryData] = useState([]);

  useEffect(() => {
    if (isConnected) {
      setLoading(true);
      client
        .query({
          query: gql`
          query MyQuery {
            amountBorroweds(where: {userAddress: "${address}"}) {
              amount
              transactionHash
              merchent
            }
          }
            `,
        })
        .then((res) => {
          setLoading(false);
          setHistoryData(res.data.amountBorroweds)
        });
    }
  }, [isConnected]);
  return (
    <div className="w-full px-4 min-h-screen bg-black flex flex-col text-white">
      <div className="mb-4 p-4 flex justify-between">
        <Header hideClose={true} />
        {isConnected && (
          <span className="text-lg font-semibold">Hi, {address}</span>
        )}
      </div>
      <hr />
      {!isConnected && (
        <>
          <div className="flex flex-col items-center mt-9">
            <div className="items-center justify-center flex mt-6">
              <img src="wallet.png" />
              <span className="text-2xl ml-1">Connect a wallet</span>
            </div>
            <div className="mt-6   w-[300px] items-center justify-center">
              <ConnectWallet nextStage={() => {}} />
            </div>
          </div>
        </>
      )}
      {isConnected && (
        <div className="flex w-full mt-9 flex-col">
          <div className="flex flex-wrap">
            <div className="h-[100px] w-[200px] g1bg rounded-lg p-4 flex flex-col">
              <span className="text-white text-lg font-semibold">
                Current Debt
              </span>
              <span className="text-orange-200">
                {wieToEthStr(data?.debt || 0n)} ETH
              </span>
            </div>
            <div className="h-[100px] w-[200px] g2bg rounded-lg p-4 flex flex-col ml-4">
              <span className="text-white text-lg font-semibold">
                Credit Limit
              </span>
              <span className="text-orange-300">
              { wieToEthStr(data?.creditLimit - data?.debt || 0n)} / { wieToEthStr(data?.creditLimit || 0n)} ETH
              </span>
            </div>
            <div className="h-[100px] w-[200px] g3bg bg-white rounded-lg p-4 flex flex-col ml-4">
              <span className="text-white text-lg font-semibold">
                Pay-Dues In
              </span>
              <span className="text-orange-200">
                {30 - new Date().getDate() + 10} days
              </span>
            </div>
            <div className="h-[100px] w-[220px] g4bg rounded-lg p-4 flex flex-col ml-4">
              <span className="text-white text-lg font-semibold">Repay</span>
              <button className="bg-[gray] p-1 px-2 rounded-lg active:scale-90 text-white font-semibold ">
                {wieToEthStr(data?.debt || 0n)} ETH +{" "}
                {Number(wieToEthStr(data?.debt || 0n)) * 2} ENT
              </button>
            </div>
          </div>
          <div className="mt-8 flex flex-col">
            <span className="text-xl font-semibold mb-4">
              Transaction History:
            </span>
            {isLoading && <CircularProgress />} 
            {!isLoading && <div className="w-[1200px]">
              <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell>#ID</TableCell>
                      <TableCell align="left">Merchent</TableCell>
                      <TableCell align="left">Amount</TableCell>
                      <TableCell align="left">TxHash</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                  {txHistory.map((row, index) => (
            <TableRow
              key={index}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                #{index  + 1}
              </TableCell>
              <TableCell align="left">{row.merchent}</TableCell>
              <TableCell align="left">{wieToEthStr(row.amount)}</TableCell>
              <TableCell align="left" sx={{ color: "blue"}}><a href={`https://base-sepolia.blockscout.com/tx/${row.transactionHash}`}>{row.transactionHash}</a></TableCell>
            </TableRow>
          ))}

                  </TableBody>
                </Table>
              </TableContainer>
            </div>}
          </div>
        </div>
      )}
    </div>
  );
}
