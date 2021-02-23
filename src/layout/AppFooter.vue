<template>
    <footer class="footer has-cards">
        <div class="card px-5 mb-5" style="margin-left: 20%; margin-right: 20%;">
            <img src="img/home/poster.png">
        </div>
        <hr>
        <div class="container">
            <div class="row align-items-center justify-content-md-between">
                <div class="col-md-6">
                    <div class="copyright">
                        &copy; {{year}}
                        <a href="https://www.keiokoeken.com" target="_blank" rel="noopener">こえけん！アニメ声優研究会</a> 
                    </div>
                </div>
                <div class="col-md-6">
                    <ul class="nav nav-footer justify-content-end">
                        <li class="nav-item">
                            <a href="https://www.keiokoeken.com/blog/" class="nav-link" target="_blank" rel="noopener">Blog</a>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </footer>
</template>
<script>
export default {
  name: 'app-footer',
  data() {
    return {
      year: new Date().getFullYear()
    }
  }
};
</script>
<style>
</style>
