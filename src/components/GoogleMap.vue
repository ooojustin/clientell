<template>
    <div class="flex justify-center">
        <iframe :src="srcUrl" :width="width" :height="height"></iframe>
    </div>
</template>

<script>
export default {
    name: 'GoogleMap',
    props: {
        placeId: String,
        width: String,
        height: String
    },
    computed: {
        srcUrl() {
            const apiKey = "AIzaSyAeN85xIwoCXFUd3ZcSLLWVf0_LNPOm6Jo";
            return `https://www.google.com/maps/embed/v1/place?key=${apiKey}&q=place_id:${this.placeId}`;
        }
    }
}
</script>
