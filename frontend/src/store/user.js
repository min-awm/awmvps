import API from "@/api";
import axios from "@/axios";
import { defineStore } from "pinia";
import router from "@/router";

const DEFINE_USER_INFO = {
  username: null,
};

const defaultState = () => {
  return {
    userInfo: DEFINE_USER_INFO,
    balance: 0,
    settings: {
      minDeposit: null,
      minWithdraw: null,
      bank: {
        bankName: null,
        accountNumber: null,
        holder: null,
        image: null,
      },
    },
    isLogin: false,
    showLoginMoblie: 0,
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
          this.userInfo = res.data;
          this.isLogin = true;
        } else {
          this.userInfo = DEFINE_USER_INFO;
        }
      } catch (error) {
        console.log(error);
      }
    },
    
    updateUserInfo(payload) {
      Object.assign(this.userInfo, payload);
    },

    logout() {
      localStorage.removeItem("accessToken");
      this.$reset();
      router.push("/");
    },

    setLogin(value) {
      this.isLogin = value;
    },

    setShowLoginMoblie() {
      this.showLoginMoblie += 1;
    },

    $reset() {
      Object.assign(this, defaultState());
    },
  },
});
