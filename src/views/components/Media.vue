<template>
<div class="bg-warning">
    <div class="container container-lg pb-5">
        <div class="h2 pl-4 py-4">メディア</div>
        <div class="row">
            <div class="col-md-4 mb-5 mb-md-0">
                <a href="https://twitter.com/keiokoeken">
                    <div class="card card-lift--hover shadow border-0">
                        <i class="pt-3 pl-3 pb-3 fa fa-2x fa-twitter"></i>
                            <img v-lazy="'img/home/koeken-twitter.png'" class="card-img">
                        <div class="pt-5 px-2 text-center">
                            <p>Twitter ではイベントの情報やサークル員のブログの更新情報などをツイートしています！</p>
                        </div>
                    </div>
                </a>
            </div>
            <div class="col-md-4 mb-5 mb-lg-0">
                <a href="https://www.youtube.com/channel/UC206slBBxrf-pPY6hHVox5Q">
                    <div class="card card-lift--hover shadow border-0">
                        <i class="pt-3 pl-3 pb-3 fa fa-2x fa-youtube"></i>
                            <img v-lazy="'img/home/koeken-youtube.png'" class="card-img">
                        <div class="pt-5 px-2 text-center">
                            <p>Youtube ではサークル員のゲーム動画や新入生向けの動画を投稿しています！</p>
                        </div>
                    </div>
                </a>
            </div>
            <div class="col-md-4 mb-5 mb-md-0">
                <a href="https://www.keiokoeken.com/blog/">
                    <div class="card card-lift--hover shadow border-0">
                        <i class="pt-3 pl-3 pb-3 fa fa-2x fa-pencil"></i>
                            <img v-lazy="'img/home/koeken-blog.png'" class="card-img">
                        <div class="pt-5 px-2 text-center">
                            <p>こえけん！に所属するサークルメンバーたちによるブログを公開しています！</p>
                        </div>
                    </div>
                </a>
            </div>
        </div>
    </div>
</div>
</template>

