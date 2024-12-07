import { createRouter, createWebHashHistory } from "vue-router";
import { useUserStore } from "@/store/user";
import MainLayout from "@/components/layouts/MainLayout.vue";

const routes = [
  {
    path: "/login",
    name: "login",
    component: () => import("@/components/auth/Login.vue"),
    meta: { redirectIfLoggedIn: true },
  },

  {
    path: "/",
    name: "index",
    component: MainLayout,
    meta: { auth: true },
    redirect: { path: "/home" },
    children: [
      {
        path: "/home",
        name: "home",
        component: () => import("@/components/home/HomePage.vue"),
      },
      {
        path: "/files",
        name: "files",
        component: () => import("@/components/files/FileManager.vue"),
      },
      {
        path: "/terminal",
        name: "terminal",
        component: () => import("@/components/terminal/TerminalPage.vue"),
      },
      {
        path: "/databases",
        children: [
          {
            path: "",
            name: "databases-page",
            component: () => import("@/components/databases/DatabasePage.vue"),
          },
          {
            path: "mariadb",
            name: "databases-mariadb",
            component: () => import("@/components/databases/Mariadb.vue"),
          },
        ],
      },
      {
        path: "/docker",
        name: "docker",
        component: () => import("@/components/docker/DockerPage.vue"),
        redirect: { name: "container" },
        children: [
          {
            path: "container",
            name: "container",
            component: () => import("@/components/docker/Container.vue"),
          },
          {
            path: "image",
            name: "image",
            component: () => import("@/components/docker/Image.vue"),
          },
          {
            path: "volume",
            name: "volume",
            component: () => import("@/components/docker/Volume.vue"),
          },
          {
            path: "setting",
            name: "setting",
            component: () => import("@/components/docker/Setting.vue"),
          },
        ],
      },
      {
        path: "/nginx",
        name: "nginx",
        component: () => import("@/components/nginx/NginxPage.vue"),
      },
      {
        path: "/security",
        children: [
          {
            path: "",
            name: "security-page",
            component: () => import("@/components/security/SecurityPage.vue"),
          },
          {
            path: "block-ip",
            name: "block-ip",
            component: () => import("@/components/security/BlockIp.vue"),
          },
          {
            path: "port",
            name: "port",
            component: () => import("@/components/security/Port.vue"),
          },
          {
            path: "ssl",
            name: "ssl",
            component: () => import("@/components/security/Ssl.vue"),
          },
        ],
      },
    ],
  },

  {
    path: "/mini-terminal",
    name: "mini-terminal",
    component: () => import("@/components/terminal/Xterm.vue"),
    meta: { auth: true },
  },

  {
    path: "/a",
    name: "a",
    component: () => import("@/components/terminal/a.vue"),
    meta: { auth: true },
  },

  {
    path: "/:pathMatch(.*)*",
    redirect: { name: "home" },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

router.beforeEach((to, _, next) => {
  const isLoggedIn = localStorage.getItem("accessToken");
  const userStore = useUserStore();

  if (to.meta.auth && !isLoggedIn) {
    return next({ name: "login" });
  }

  if (to.meta.redirectIfLoggedIn && isLoggedIn) {
    return next({ name: "home" });
  }

  if (to.meta.auth && isLoggedIn) {
    // Get userInfo when route change
    userStore.getUserInfo();
  }

  return next();
});

export default router;
