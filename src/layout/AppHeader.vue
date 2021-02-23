<template>
    <header class="header-global sticky-top">
        <base-nav class="navbar-main header-color" transparent type="" effect="light" expand>
            <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
                <img src="img/brand/koeken_logo.png" alt="logo">
            </router-link>

            <div class="row" slot="content-header" slot-scope="{closeMenu}">
                <div class="col-6 collapse-brand">
                    <router-link to="/" >
                        <img src="img/brand/koeken_logo.png">
                    </router-link>
                </div>
                <div class="col-6 collapse-close">
                    <close-button @click="closeMenu"></close-button>
                </div>
            </div>

            <ul class="navbar-nav navbar-nav-hover align-items-lg-center">
                <base-dropdown class="nav-item" menu-classes="dropdown-menu-xl">
                    <a slot="title" href="#" class="nav-link" data-toggle="dropdown" role="button">
                        <i class="ni ni-ui-04 d-lg-none"></i>
                        <span class="nav-link-inner--text">Pages</span>
                    </a>
                    <div class="dropdown-menu-inner">
                        <router-link to="/about" class="media d-flex align-items-center">
                            <div class="icon icon-shape bg-gradient-primary rounded-circle text-white">
                                <i class="fa fa-at"></i>
                            </div>
                            <div class="media-body ml-3">
                                <h6 class="heading text-primary mb-md-1">About Us</h6>
                                <p class="description d-none d-md-inline-block mb-0">こえけん！アニメ声優研究会の活動について</p>
                            </div>
                        </router-link>
                        <router-link to="/events" class="media d-flex align-items-center">
                            <div class="icon icon-shape bg-gradient-warning rounded-circle text-white">
                                <i class="fa fa-child"></i>
                            </div>
                            <div class="media-body ml-3">
                                <h5 class="heading text-warning mb-md-1">Events</h5>
                                <p class="description d-none d-md-inline-block mb-0">今までこえけん！で行われたイベントについて</p>
                            </div>
                        </router-link>
                        <router-link to="/activity" class="media d-flex align-items-center">
                            <div class="icon icon-shape bg-gradient-success rounded-circle text-white">
                                <i class="fa fa-tag"></i>
                            </div>
                            <div class="media-body ml-3">
                                <h5 class="heading text-success mb-md-1">Activities</h5>
                                <p class="description d-none d-md-inline-block mb-0">サークル内での活動履歴</p>
                            </div>
                        </router-link>
                        <router-link to="/newcomers" class="media d-flex align-items-center">
                            <div class="icon icon-shape bg-gradient-info rounded-circle text-white">
                                <i class="fa fa-thumb-tack"></i>
                            </div>
                            <div class="media-body ml-3">
                                <h5 class="heading text-info mb-md-1">New Comers</h5>
                                <p class="description d-none d-md-inline-block mb-0">入会を希望する新入生の方へ</p>
                            </div>
                        </router-link>
                    </div>
                </base-dropdown>
            </ul>
            <ul class="navbar-nav align-items-lg-center ml-lg-auto">

                <li class="nav-item">
                    <a class="nav-link nav-link-icon" href="https://twitter.com/keiokoeken" target="_blank" rel="noopener"
                       data-toggle="tooltip" title="Follow us on Twitter">
                        <i class="fa fa-twitter-square"></i>
                        <span class="nav-link-inner--text d-lg-none">Twitter</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link nav-link-icon" href="https://github.com/creativetimofficial/vue-argon-design-system"
                       target="_blank" rel="noopener" data-toggle="tooltip" title="Star us on Github">
                        <i class="fa fa-youtube"></i>
                        <span class="nav-link-inner--text d-lg-none">Youtube Channel</span>
                    </a>
                </li>
                <li class="nav-item d-none d-lg-block ml-lg-4">
                    <a href="#contact"
                       class="btn btn-neutral btn-icon">
                <span class="btn-inner--icon">
                  <i class="fa fa-envelope-o mr-2"></i>
                </span>
                        <span class="nav-link-inner--text">お問い合わせ</span>
                    </a>
                </li>
            </ul>
        </base-nav>
    </header>
</template>
<script>
import BaseNav from "@/components/BaseNav";
import BaseDropdown from "@/components/BaseDropdown";
import CloseButton from "@/components/CloseButton";

export default {
  components: {
    BaseNav,
    CloseButton,
    BaseDropdown
  }
};
</script>
<style>
</style>
