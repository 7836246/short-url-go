// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  //运行时配置
  runtimeConfig: {
    public: {
      siteApi: process.env.SITE_API,
    },
  },
  // 应用配置
  app: {
    head: {
      charset: 'utf-8',
      htmlAttrs: {
        lang: 'zh-CN'
      },
      meta: [
        {name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no'},
      ],
      link: [
      ],
      script: [
      ]
    }
  },
  css: ["assets/css/style.css"],
  modules: [
    '@nuxt/devtools',
    'arco-design-nuxt-module',
    '@unocss/nuxt',
    '@nuxtjs/color-mode',
    "@nuxt/icon"
  ],
  imports: {
    autoImport: true,
  },
  arco: {
    importPrefix: 'A',
    hookPrefix: 'Arco',
    locales: ['getLocale'],
    localePrefix: 'Arco',
  },
  //打包配置
  nitro: {
    //缩小捆绑包
    minify: true,
    //关闭源映射生成
    sourceMap: false,
  },
  vite: {
    build: {
      chunkSizeWarningLimit: 1500,
    },
  },
  //取消取消行内样式
  features: {
    inlineStyles: false,
  },
  // @ts-ignore
  meta: {
    renderSSRHeadOptions: { // typeof RenderSSRHeadOptions @unhead/schema
      omitLineBreaks: true,
    }
  },
  future: {
    compatibilityVersion: 4,
  },
})
