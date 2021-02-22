<template>
    <section id="contact" class="section pb-0 section-components bg-info">
        <form v-on:submit.prevent="postContactInfo">
            <div class="container mb-5">
                <!-- Inputs -->
                <h3 class="h2 text-body font-weight-bold mb-4">お問い合わせフォーム</h3>
                <p>サークル員が対応いたします。</p>
                <div class="mb-3">
                    <small class="text-body font-weight-bold">お名前</small>
                    <p class="text-danger"> {{ name_error }} </p>
                </div>
                <div class="row">
                    <div class="col-lg-4 col-sm-6">
                        <base-input v-model="name" name="name" value="name" placeholder="こえけん 太郎" :required="true"></base-input>
                    </div>
                </div>
                <div class="mb-3">
                    <small class="text-body font-weight-bold">メールアドレス</small>
                    <p class="text-danger"> {{ email_error }} </p>
                </div>
                <div class="row">
                    <div class="col-lg-4 col-sm-6">
                        <base-input v-model="email" name="email" value="email" placeholder="koeken.taro@gmail.com" :required="true"></base-input>
                    </div>
                </div>
                <div class="mb-3">
                    <small class="text-body font-weight-bold">お問い合わせ内容</small>
                    <p class="text-danger"> {{ text_error }} </p>
                </div>
                <div class="row px-3">
                    <textarea v-model="text" name="text" value="text" id="text" class="form-control form-control-alternative mb-5" rows="5" placeholder="Write your questions here..."></textarea>
                </div>
                <div class="pb-5">
                    <base-button nativeType="submit" round block size="lg" v-if="send_status">送信する</base-button>
                    <p class="text-dark" v-else>お問い合わせありがとうございます！</p>
                </div>
            </div>
        </form>
    </section>
</template>
<script>
import axios from 'axios';
export default {
    data: function() {
        return {
            errors: [],
            name: '',
            email: '',
            text: '',
            name_error: '',
            email_error: '',
            text_error: '',
            send_status: true
        }
    },
    methods: {
        postContactInfo() {
            var params = new URLSearchParams();
            params.append("name", this.name);
            params.append("email", this.email);
            params.append("text", this.text);
            let that = this;
            const url = 'http://127.0.0.1:8081/api/contactForm';
            axios.post(url, params)
            .then(res => {
                console.log(res);
                console.log(Object.keys(res.data).length);
                if(Object.keys(res.data).length > 0){
                    that.name_error = res.data['Name'];
                    that.email_error = res.data['Email'];
                    that.text_error = res.data['Text'];
                } else {
                    that.send_status = false;
                    that.name_error = '';
                    that.email_error = '';
                    that.text_error = '';
                }
            })
            .catch(
                err => {
                    console.log('cannot send')
                }
            );
        }
    }
}
</script>