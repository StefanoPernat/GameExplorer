<template>
  <div class="home">
    Home {{ keyWatcher }}
  </div>
</template>

<script>

import { mapState } from 'vuex'

export default {
  data () {
    return {}
  },
  computed: {
    ...mapState([
      'offers'
    ]),
    keyWatcher: function () {
      console.log('aaa')
      return process.env.VUE_APP_PREVIEW_KEY
    }
  }

}
</script>
