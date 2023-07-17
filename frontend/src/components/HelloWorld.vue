<template>
  <div><img v-bind:src="gifPath" alt=""></div>
</template>

<script>
export default {
   data(){
    return{
      gifPath:"./src/assets/images/draw.gif"
    }
  }
}
</script>