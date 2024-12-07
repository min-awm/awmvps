import API from "@/api";
import axios from "@/axios";
import { defineStore } from "pinia";
import router from "@/router";

const DEFINE_USER_INFO = {
  name: null,
};

const defaultState = () => {
  return {
    userInfo: DEFINE_USER_INFO,  
    isLogin: false,
  };
};

export const useUserStore = defineStore("user", {
  state: () => defaultState(),

  getters: {},

  actions: {
    async getUserInfo() {
      try {
        const res = await axios.get(API.USER_INFO);
        if (res.success) {
          this.userInfo = { name: res.name };
          this.isLogin = true;
        } else {
          this.userInfo = DEFINE_USER_INFO;
        }
      } catch (error) {
        console.log(`Store User: ${error}`);
      }
    },

    logout() {
      localStorage.removeItem("accessToken");
      this.$reset();
      router.push({ name: "login" });
    },

    $reset() {
      Object.assign(this, defaultState());
    },
  },
});
