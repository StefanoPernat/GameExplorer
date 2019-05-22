<template>
  <div class="home">
    Home {{ keyWatcher }}
  </div>
</template>

<script>

import { mapState } from 'vuex'
import { mapActions } from 'vuex'

export default {
  data () {
    return {}
  },
  computed: {
    ...mapState([
      'offers'
    ]),
    ...mapActions([
      'fetchDeals'
    ]),
    keyWatcher: function () {
      return process.env.VUE_APP_PREVIEW_KEY
    }
  },
  mounted () {
    this.fetchDeals()
  }

}
</script>
