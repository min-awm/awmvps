import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "@/store/user";
import MainLayout from "@/components/layouts/MainLayout.vue";

const routes = [
  {
    path: "/",
    name: "index",
    component: MainLayout,
    children: [
      {
        path: "/",
        name: "home",
        component: () => import("@/components/home/HomePage.vue"),
      },
      {
        path: "/files",
        name: "files",
        component: () => import("@/components/files/FilePage.vue"),
      },
      {
        path: "/security",
        name: "security",
        component: () => import("@/components/security/SecurityPage.vue"),
      },
    ],
  },

  {
    path: "/:pathMatch(.*)*",
    redirect: { name: "home" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

// router.beforeEach((to, _, next) => {
//   const isLoggedIn = localStorage.getItem("accessToken");
//   const userStore = useUserStore();

//   if (to.meta.auth && !isLoggedIn) {
//     userStore.setLogin(false);
//     userStore.setShowLoginMoblie();
//     return next({ name: "home" });
//   }

//   if (to.meta.redirectIfLoggedIn && isLoggedIn) {
//     return next({ name: "home" });
//   }

//   if (to.meta.auth && isLoggedIn) {
//     // Get userInfo when route change
//     userStore.getUserInfo();
//   }

//   return next();
// });

export default router;
