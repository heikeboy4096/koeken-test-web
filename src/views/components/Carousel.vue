<template>
    <section class="section section-lg section-shaped overflow-hidden my-0">
        <div class="shape shape-style-1 shape-default shape-skew">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
        </div>
        <div class="container my-5 pb-5 pb-lg">
            <div class="row justify-content-between align-items-center">
                <div class="col-lg-5 mb-5 mb-lg-0">
                    <h1 class="text-info font-weight-light">声優さんをお呼びしてイベントを開催！</h1>
                    <p class="lead text-white mt-4">こえけん！アニメ声優研究会では有名な声優さんを三田祭や矢上祭などの学園祭にお呼びしてトークショー等を開催しています！</p>
                    <a href="/event"
                       class="btn btn-white mt-4">イベントを見る</a>
                </div>
                <div class="col-lg-6">
                    <div class="rounded shadow-lg overflow-hidden bg-white transform-perspective-right">
                        <b-carousel id="carousel1"
                                    controls
                                    indicators>
                            <!-- Text slides with image -->
                            <b-carousel-slide img-src="img/events/2019mita.png"></b-carousel-slide>
                            <b-carousel-slide img-src="img/events/2017yagami1.png"></b-carousel-slide>
                            <b-carousel-slide img-src="img/events/2017yagami2.png"></b-carousel-slide>
                        </b-carousel>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
<script>
import { BCarousel } from "bootstrap-vue/esm/components/carousel/carousel";
import { BCarouselSlide } from "bootstrap-vue/esm/components/carousel/carousel-slide";

export default {
  components: {
    BCarousel,
    BCarouselSlide
  }
};
</script>
<style>
</style>
