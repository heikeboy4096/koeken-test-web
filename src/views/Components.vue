<template>
    <div>
        <hero></hero>
        <carousel></carousel>
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
import BasicElements from "./components/BasicElements";
import Inputs from "./components/Inputs";
import CustomControls from "./components/CustomControls";
import Navigation from "./components/Navigation";
import JavascriptComponents from "./components/JavascriptComponents";
import Carousel from "./components/Carousel";
import Icons from "./components/Icons";
import Examples from "./components/Examples";
import DownloadSection from "./components/DownloadSection";
import ContactForm from "./components/ContactForm";
import Media from "./components/Media";

export default {
  name: "components",
  components: {
    Hero,
    BasicElements,
    Inputs,
    CustomControls,
    Navigation,
    JavascriptComponents,
    Carousel,
    Icons,
    Examples,
    DownloadSection,
    //added by me
    ContactForm,
    Media,
  }
};
</script>
