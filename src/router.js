import Vue from "vue";
import Router from "vue-router";
import AppHeader from "./layout/AppHeader";
import AppFooter from "./layout/AppFooter";
import Home from "./views/Home.vue";
//import Landing from "./views/Landing.vue";
//import Login from "./views/Login.vue";
//import Register from "./views/Register.vue";
//import Profile from "./views/Profile.vue";
import About from "./views/About.vue";
import Events from "./views/Events.vue";
import Activity from "./views/Activity.vue";
import Newcomers from "./views/Newcomers.vue";

Vue.use(Router);

export default new Router({
  mode: 'history',
  linkExactActiveClass: "active",
  routes: [
    {
      path: "/",
      name: "home",
      components: {
        header: AppHeader,
        default: Home,
        footer: AppFooter
      }
    },
    {
      path: "/about",
      name: "About",
      components: {
        header: AppHeader,
        default: About,
        footer: AppFooter
      }
    },
    {
      path: "/events",
      name: "Events",
      components: {
        header: AppHeader,
        default: Events,
        footer: AppFooter
      }
    },
    {
      path: "/activity",
      name: "Activity",
      components: {
        header: AppHeader,
        default: Activity,
        footer: AppFooter
      }
    },
    {
      path: "/newcomers",
      name: "Newcomers",
      components: {
        header: AppHeader,
        default: Newcomers,
        footer: AppFooter
      }
    }
  ],
  scrollBehavior: to => {
    if (to.hash) {
      return { selector: to.hash };
    } else {
      return { x: 0, y: 0 };
    }
  }
});
