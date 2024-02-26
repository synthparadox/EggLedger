<template>
    <button
        title="Remove filter condition" type="button" :id="internalId"
        class="mr-1rem flex items-center justify-center text-red-700 text-lg max-w-6 max-h-6 min-w-6 min-h-6 border border-red-700 rounded-md bg-transparent mt-05rem filter-button"
        v-on:click="e => { e.preventDefault(); $emit('removeAndShift', index, orIndex)}"
    >
    ×
    </button>
</template>

<script>
    export default {
        emits: ['removeAndShift'],
        props: {
            disabledBool: Boolean,
            index: Number,
            orIndex: Number,
            internalId: String,
        },
    }
</script>