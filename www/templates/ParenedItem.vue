<template>
    <span :class="getParenClass()">(<span :class="innerClass">{{ getInner() }}</span>)</span>
</template>

<script>
    export default {
        props: {
            parenClass: String,
            innerClass: String,
            innerText: String,
            includeSpaces: Boolean,
        },
        methods: {
            getInner(){
                return this.includeSpaces ? ` ${this.innerText} ` : this.innerText;
            },
            getParenClass(){
                return this.parenClass ?? 'text-gray-400';
            }
        }
    }
</script>