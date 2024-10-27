import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

interface SpeechBubbleProps {
  description: string
}

export function SpeechBubble({ description }: SpeechBubbleProps) {
  return (
    <Card className="bg-white rounded-xl border-none">
      <CardContent className="p-1">
        <p className="text-sm break-words text-gray-600 mb-2">{description}</p>
        <Button className="w-full bg-[#f87171] hover:bg-[#ef4444] text-white">
          よせがきを見る
        </Button>
      </CardContent>
    </Card>
  )
}
