import router from "@/router";
import axios from "axios";

const axiosIns = axios.create({
  baseURL: `${import.meta.env.VITE_APP_BASE_URL}/api`,
  // timeout: 1000,
  // headers: {'X-Custom-Header': 'foobar'}
});

axiosIns.interceptors.request.use((config) => {
  const token = localStorage.getItem("accessToken");
  if (token) {
    config.headers = config.headers || {};
    config.headers.Authorization = token ? `Bearer ${token}` : "";
  }
  return config;
});

axiosIns.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    console.log(error)
    if (error.response.status === 401) {
      localStorage.removeItem("accessToken");

      // If 401
      router.push("/login");
    } else {
      return Promise.reject(error.response.data);
    }
  }
);
export default axiosIns;
