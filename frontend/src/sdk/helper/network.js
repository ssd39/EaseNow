import { coreGateway } from "../constants";
import { toastError } from "./toast-helper"
import { ethers } from "ethers";

const api =  coreGateway

const registerStart = async (phone_number, country_code) => {
    const res = await (await fetch(api + "/register-start", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            phone_number,
            country_code: Number(country_code)
        })
    })).json()
    if(res.success){
        return res.session_id
    }
    toastError(res.message)
    return ""
}

const registerValidate = async (session_id, otp) => {
    const res = await (await fetch(api + "/register-validate", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            session_id,
            otp: Number(otp)
        })
    })).json()
    if(res.success){
        return wieToEthStr(res.credit_limit)
    }
    toastError(res.message)
    return ""
}


const registerFinalise = async (message, signature, ipfsHash) => {
    const res = await (await fetch(api + "/register-finalise", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            private_data: ipfsHash,
            wallet_message: message,
            wallet_sig: signature
        })
    })).json()
    if(res.success){
        return res.tx_hash
    }
    toastError(res.message)
    return ""
}

function wieToEthStr(wieNum) {
    const res = ethers.formatEther(wieNum)
    const parts = res.split(".")
    return parts[0] + "." + parts[1].slice(0, 3)
}

export {
    registerStart,
    registerValidate,
    registerFinalise,
    wieToEthStr
}