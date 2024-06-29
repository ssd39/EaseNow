import { toast } from "react-toastify";

function toastError(message) {
  toast.error(message, {
    position: "top-right",
    autoClose: 3000,
    closeOnClick: true,
    pauseOnHover: true,
    theme: "dark",
  });
}

function toastSucess(message) {
  toast.success(message, {
    position: "top-right",
    autoClose: 3000,
    closeOnClick: true,
    pauseOnHover: true,
    theme: "dark",
  });
}

export {
    toastError,
    toastSucess
}