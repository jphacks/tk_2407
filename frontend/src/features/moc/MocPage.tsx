import { Button } from '@/components/ui/button'
import React from 'react'

export default function MocPage() {
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
