import router from "@/router";
import axios from "axios";

const hostname = window.location.hostname; 
const port = window.location.port;
const ipHost = `${hostname}${port ? `:${port}` : ''}`;

const axiosIns = axios.create({
  baseURL: `http://${
    import.meta.env.VITE_APP_MODE == "dev"
      ? import.meta.env.VITE_APP_BASE_URL
      : ipHost
  }`,
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
    console.log(`Axios: ${error}`);
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
