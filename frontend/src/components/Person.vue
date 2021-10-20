<template>
    <div>
        {{ data.firstName }} {{ data.lastName }}
    </div>
</template>

<script>
export default {
    name: 'Person',
    props: {
        data: Object
    }
}
</script>

<style scoped>
</style>
