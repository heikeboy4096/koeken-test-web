<template>
    <div class="row py-5">
        <div class="mx-5 col-sm-4 card shadow order-1">
            <Timeline :id="user_id" sourceType="profile" :options="{ tweetLimit: '1' }"></Timeline>
        </div>
        <div class="col-sm-6 pt-5 pl-5 order-0">
            <div>
                <img src="img/brand/icons8-twitter.svg" class="float-left">
                <h3 style="white-space: nowrap;" class="text-primary font-weight-light">&nbsp;&nbsp;新入生向けTwitter</h3>
            </div>
            <p class="lead mt-4">まだサークルが決まっていない皆さん！</p>
            <h4 class="text-danger font-weight-bold">オタク仲間とアニメ・声優の話、したくない？</h4>
            <p class="lead mt-4">サークルに在籍するオタクの紹介や新歓に関する情報を専用のTwitterアカウントで発信していきます！</p>
        </div>
    </div>
</template>

<script>
import { Timeline } from "vue-tweet-embed";

export default {
    components: {
        Timeline
    },
    data: function() {
        return {
            user_id: "koeken_orien21"
        }
    }
}
</script>