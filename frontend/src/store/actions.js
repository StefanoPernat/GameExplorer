import axios from 'axios'

export default {
  fetchDeals: ({ commit }) => {
    axios.get('http://localhost:8081/top').then(function (response) {
      console.log(response)
      if(response.data !== null || response.data !== undefined ) {
        commit('appendTop', response.data)
      }
    }).catch(function (error) {
      console.log('Error: '+ error)
    })
  }
}
