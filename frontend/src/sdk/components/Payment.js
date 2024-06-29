import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchUser } from "../redux/actionSlice";
import { useAccount } from "wagmi";
import { CircularProgress } from "@mui/material";

export default function Payment({ nextStage }) {
  const { address } = useAccount();
  const isNewUser = useSelector((state) => state.action.isNewUser);
  const isUserFetched = useSelector((state) => state.action.isUserFetched);
  const [loading, setLoading] = useState(true);
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchUser(address));
  }, []);

  useEffect(() => {
    if (isUserFetched) {
      if (isNewUser) {
        nextStage(2);
      }
      setLoading(false);
    }
  }, [isUserFetched]);

  return (
    <div className="flex flex-1">
      {loading && (
        <div className="text-blue-600 flex items-center justify-center flex-1">
          <CircularProgress size={64} />
        </div>
      )}
    </div>
  );
}
