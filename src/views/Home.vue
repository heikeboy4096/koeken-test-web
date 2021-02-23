<template>
    <div>
        <hero></hero>
        <carousel></carousel>
        <Tweets></Tweets>
        <Media></Media>
        <ContactForm></ContactForm>
        <!--  below is template components
        <basic-elements></basic-elements>
        <inputs></inputs> 
        <custom-controls></custom-controls>
        <navigation></navigation>
        <javascript-components></javascript-components>
        <icons></icons>
        <examples></examples>
        <download-section></download-section>
        -->
    </div>
</template>

<script>
import Hero from "./components/Hero";
import Carousel from "./components/Carousel";
import ContactForm from "./components/ContactForm";
import Media from "./components/Media";
import Tweets from "./components/Tweets";

export default {
  name: "components",
  components: {
    Hero,
    Carousel,
    //added by me
    Tweets,
    ContactForm,
    Media,
  }
};
</script>
