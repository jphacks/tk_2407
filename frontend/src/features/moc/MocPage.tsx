'use client'

import { Button } from '@/components/ui/button'

import React, {useEffect} from 'react'
import api from "@/api/$api";
import aspida from '@aspida/fetch';

// to Rare-san

// 1. pnpm i @aspida/fetch
// 2. apiClientを作成（別ファイルでやって他からimportするかたちがいいかも）
//      import api from "@/api/$api";
//      import aspida from '@aspida/fetch';
//      const apiClient = api(aspida(undefined, {baseURL: 'http://localhost:8080/'}))
// 3. Codeを参考に使い方をみてもらえれば3150


export default function MocPage() {
    const apiClient = api(aspida(undefined, {baseURL: 'http://localhost:8080/'}))
  useEffect(() => {
      const fechData = async () => {
          const res = await apiClient.api.v1.health.get()
          console.log(res)
      }
      fechData()
  }, [])
  return (
    <>
      <h1>
        Moc Pageです。page.tsxには書かず、このfeaturesに書くことが多いらしい。
        apiなどのロジックをpageに出したくないとかだったような。
      </h1>
      <p>Hello World</p>
      <Button>ログイン</Button>
    </>
  )
}
