import React from "react";
import HighlightOffIcon from "@mui/icons-material/HighlightOff";
export default function Header({ onClose }) {
  return (
    <div className="flex">
      <div className="flex flex-1 flex-col select-none">
        <div>
          <span className="title text-3xl font-bold">EaseNow⚡</span>
        </div>
        <span className="text-sm text-orange-300">Buy Now Pay Later</span>
      </div>
      <HighlightOffIcon onClick={onClose} className="cursor-pointer active:scale-90" />
    </div>
  );
}
