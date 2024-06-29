import React, { useRef, useState } from "react";
import TextField from "@mui/material/TextField";
import { MenuItem } from "@mui/material";
import GButton from "./GButton";
import {
  registerFinalise,
  registerStart,
  registerValidate,
} from "../helper/network";
import { toastError, toastSucess } from "../helper/toast-helper";
import OtpInput from "react-otp-input";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import FingerprintIcon from "@mui/icons-material/Fingerprint";
import { handleUpload } from "../helper/ipfs";
import { useAccount, useSignMessage } from "wagmi";
import { SiweMessage } from "siwe";

export default function Registration() {
  const [step, setStep] = useState(0);
  const [sessionId, setSessionId] = useState("");
  const [countryCode, setCountryCode] = useState("91");
  const [phoneNumner, setPhoneNumber] = useState("");
  const [loading, setLoading] = useState(false);
  const [otp, setOtp] = useState("");
  const [creditLimit, setCreditLimit] = useState("");
  const [isUploaded, setUploaded] = useState(false);
  const [isFacialUploaded, setFacialUploaded] = useState(false);
  const [isUploading, setUploading] = useState(true);
  const [ipfsHash, setIpfsHash] = useState("");
  const [txHash, setTxHash] = useState("");

  const fileInputRef = useRef(null);
  const { address, chainId } = useAccount();

  const { signMessageAsync } = useSignMessage();

  const handleFileChange = async (event) => {
    setUploaded(true);
    const hash = await handleUpload(event.target.files[0]);
    setUploading(false);
    setIpfsHash(hash);
    if (hash == "") {
      setUploaded(false);
      setUploading(true);
    }
  };

  const step1 = async () => {
    setLoading(true);
    if (phoneNumner != "") {
      await registerStart(phoneNumner, countryCode)
        .then((sessionId) => {
          if (sessionId != "") {
            setSessionId(sessionId);
            setStep(1);
          }
        })
        .catch((e) => toastError("Network error!"));
    } else {
      toastError("Phone number is missing!");
    }
    setLoading(false);
  };

  const step2 = async () => {
    setLoading(true);
    if (otp != "") {
      await registerValidate(sessionId, otp)
        .then((creditLimit) => {
          if (creditLimit != "") {
            setCreditLimit(creditLimit);
            setStep(2);
          }
        })
        .catch((e) => toastError("Network error!"));
    } else {
      toastError("Please enter the otp!");
    }
    setLoading(false);
  };

  const step3 = async (message) => {
    setLoading(true);
    if (ipfsHash != "") {
      try {
        const signature = await signMessageAsync({
          message: message.prepareMessage(),
        }).catch((e) => toastError("Failed to get signature!"));
        if (signature != "") {
          await registerFinalise(message.prepareMessage(), signature, ipfsHash).then(
            (txHash) => {
              if (txHash != "") {
                setTxHash(txHash);
                setStep(3);
              }
            }
          );
        }
      } catch (e) {
        toastError("Network error!");
      }
    } else {
      toastError("Please complete required steps to continue!");
    }
    setLoading(false);
  };

  return (
    <>
      {step == 0 && (
        <div className="mt-6 flex flex-col items-center w-full ">
          <span className="text-lg">Hi, Welcome to EaseNow world!</span>
          <span className="text-lg text-center">
            We are here to make your{" "}
            <span className="text-[#3F5EFB] font-semibold">Shopping</span>/
            <span className="text-[#FC466B] font-semibold">Purchase</span>{" "}
            experience better 💫
          </span>
          <div className="mt-8">
            <TextField
              sx={{
                background: "#383838",
                color: "white",
                borderRadius: 4,
                width: 80,
                input: { color: "white" },
                ".MuiSelect-select ": { color: "rgb(253, 186, 116)" },
                ".MuiSvgIcon-root": { color: "white" },
              }}
              id="filled-basic"
              variant="filled"
              InputProps={{ disableUnderline: true }}
              select
              hiddenLabel
              defaultValue={countryCode}
              onChange={(e) => setCountryCode(e.target.value)}
            >
              <MenuItem key={"1"} value={"1"}>
                +1
              </MenuItem>
              <MenuItem key={"91"} value={"91"}>
                +91
              </MenuItem>
            </TextField>
            <TextField
              sx={{
                background: "#383838",
                color: "white",
                borderRadius: 4,
                marginLeft: 1,
                input: { color: "white" },
                label: { color: "rgb(253, 186, 116)" },
                "label.Mui-focused": { color: "rgb(253, 186, 116)" },
              }}
              id="filled-basic"
              label="Phone Number"
              variant="filled"
              value={phoneNumner}
              onChange={(e) => setPhoneNumber(e.target.value)}
              InputProps={{ disableUnderline: true }}
            />
          </div>
          <div className="flex items-center mt-4">
            <span className="text-[10.5px] text-center">
              We are using your phone number to get your credit score in
              decentralised verifiable & trusted enviorement
            </span>
          </div>
          <div className="flex-1 items-end flex mb-6">
            <GButton
              loading={loading}
              onClick={() => {
                step1();
              }}
            >
              Next
            </GButton>
          </div>
        </div>
      )}
      {step == 1 && (
        <>
          <div className="mt-6 flex flex-col items-center w-full">
            <div className="flex-1 flex-col flex items-center justify-center">
              <span className="text-lg text-center">
                We have just delivered{" "}
                <span className="text-[#FC466B] font-semibold">OTP</span> to
                you!
              </span>
              <OtpInput
                value={otp}
                onChange={setOtp}
                numInputs={6}
                renderSeparator={<span className="text-orange-300">-</span>}
                containerStyle="text-white w-full flex items-center justify-center"
                renderInput={(props) => (
                  <input
                    {...props}
                    className="bg-[#383838] m-2 py-2 !w-[2.8em] rounded-md text-white"
                  />
                )}
              />
            </div>
            <div className="mb-6">
              <GButton
                loading={loading}
                onClick={() => {
                  step2();
                }}
              >
                Next
              </GButton>
            </div>
          </div>
        </>
      )}
      {step == 2 && (
        <div className="mt-4 flex flex-col items-center w-full">
          <span className="text-lg text-center">
            Congratulations🥳, you received the credit limit of
            <span className="text-[#3F5EFB] font-semibold"> {creditLimit}</span>
            <span className="text-[#FC466B] font-semibold"> ETH!</span> Please
            complete following steps for credit safety.
          </span>
          <div className=" flex flex-col">
            <div className="flex flex-col mt-4">
              {isUploaded && (
                <div className="bg-[#0052FF] flex   px-2 rounded-t-lg">
                  <span className="text-[10px] font-semibold">
                    {isUploading ? "Uploading..." : "Completed!"}
                  </span>
                </div>
              )}
              <button
                onClick={() => {
                  if (isUploaded) {
                    return;
                  }
                  fileInputRef.current.click();
                }}
                className={`${
                  isUploaded
                    ? "bg-[gray] cursor-not-allowed"
                    : "bg-[#383838] hover:bg-[rgb(50,50,50)]  active:scale-90  rounded-lg"
                } w-[100%] flex p-3 items-center justify-center`}
              >
                Upload Personal Document <CloudUploadIcon className="ml-2" />
              </button>
              <input
                type="file"
                ref={fileInputRef}
                style={{ display: "none" }}
                onChange={handleFileChange}
              />
            </div>
            <div className="flex flex-col mt-4">
              {isFacialUploaded && (
                <div className="bg-[#0052FF] flex   px-2 rounded-t-lg">
                  <span className="text-[10px] font-semibold">Completed!</span>
                </div>
              )}
              <button
                onClick={() => setFacialUploaded(true)}
                className={`${
                  isFacialUploaded
                    ? "bg-[gray] cursor-not-allowed"
                    : "bg-[#383838] hover:bg-[rgb(50,50,50)]  active:scale-90  rounded-lg"
                } w-[100%] flex p-3 items-center justify-center`}
              >
                Facial verification <FingerprintIcon className="ml-2" />
              </button>
            </div>
          </div>
          <span className="text-[10.5px] text-center mt-1">
            We are uploading your private data on autonomus encrypted vault to
            insure re-payments and credit safety.{" "}
            <a className="text-blue-600" href="">
              Learn more.
            </a>
          </span>
          <div className="flex-1 items-end flex mb-6">
            <GButton
              loading={loading}
              onClick={async () => {
                if (!isFacialUploaded || !isUploaded) {
                  toastError(
                    "Please complete required steps before moving forward!"
                  );
                  return;
                }
                const message = new SiweMessage({
                  domain: window.location.host,
                  address,
                  statement: "Sign in with Ethereum to the app.",
                  uri: window.location.origin,
                  version: "1",
                  chainId,
                  nonce: sessionId,
                });
                step3(message);
              }}
            >
              Generate signature
            </GButton>
          </div>
        </div>
      )}
    </>
  );
}
