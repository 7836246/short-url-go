// @ts-ignore
import {joinURL} from "ufo";
// @ts-ignore
import {proxyRequest} from "h3";

export default defineEventHandler(async (event:any) =>{
    const { siteApi } = useRuntimeConfig().public
    const path = event.path.replace(/^\/api\//,'')
    const target = joinURL(siteApi,path)
    return proxyRequest(event,target)
})
