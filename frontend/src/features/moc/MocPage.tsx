'use client'

import { Button } from '@/components/ui/button'
import { apiClient } from '@/lib/aspida'

import React, { useEffect } from 'react'

// to Rare-san

// 1. pnpm i @aspida/fetch
// 2. apiClientを作成（別ファイルでやって他からimportするかたちがいいかも）
//      import api from "@/api/$api";
//      import aspida from '@aspida/fetch';
//      const apiClient = api(aspida(undefined, {baseURL: 'http://localhost:8080/'}))
// 3. Codeを参考に使い方をみてもらえれば3150

export default function MocPage() {
  useEffect(() => {
    const fetchData = async () => {
      const res = await apiClient.api.v1.health.get()
      console.log(res)
    }
    fetchData()
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
