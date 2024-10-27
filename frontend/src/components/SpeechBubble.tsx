import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import Link from 'next/link'

interface SpeechBubbleProps {
  lint: string
  description: string
}

export function SpeechBubble({ link, description }: SpeechBubbleProps) {
  return (
    <Card className="bg-white rounded-xl border-none">
      <CardContent className="p-1">
        <p className="text-sm break-words text-gray-600 mb-2">{description}</p>
        <Button className="w-full bg-[#f87171] hover:bg-[#ef4444] text-white ">
          <Link href={link}>よせがきを見る</Link>
        </Button>
      </CardContent>
    </Card>
  )
}
