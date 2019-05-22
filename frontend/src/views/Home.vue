<template>
  <div class="home">
    Home
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
    ])
  }

}
</script>
