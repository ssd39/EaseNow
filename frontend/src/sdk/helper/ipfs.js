import { toastError } from "./toast-helper";

const projectId = "ca78111a-28cc-4fdb-80f4-83758e014d09"
const projectServerKey = "smU00pB2XbveEfV2Dr3DgYU3nTPGJ2LP2aue1hfQ"
const IPFS_URL = 'https://rpc.particle.network/ipfs/upload';

const handleUpload = async (file) => {
    const formData = new FormData();
    formData.append('file', file);

    const headers = new Headers();
    headers.append('Authorization', 'Basic ' + btoa(`${projectId}:${projectServerKey}`));

    try {
        const response = await fetch(IPFS_URL, {
            method: 'POST',
            body: formData,
            headers: headers,
        });

        if (response.ok) {
            const data = await response.json();
            console.log(data);
            return data.cid
        } else {
            console.error('Upload failed:', response.statusText);
            toastError("File upload failed!")
        }
    } catch (error) {
        console.error('Error:', error);
        toastError("File upload failed!")
    }
    return ""
};

export { handleUpload }