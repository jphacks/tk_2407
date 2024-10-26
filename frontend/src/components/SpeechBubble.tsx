import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import Image from 'next/image'

interface SpeechBubbleProps {
  title: string;
  imageUrl: string;
  description: string;
}

export default function SpeechBubble({ title, imageUrl, description }: SpeechBubbleProps) {
  return (
    <div className="p-12 min-h-screen bg-gray-100">
      <div className="relative inline-block">
        <Card className="w-[400px] bg-white rounded-xl overflow-hidden border-none">
          <CardContent className="p-5">
            <div className="flex gap-4 mb-4">
              <div className="flex-shrink-0">
                <Image
                  src={imageUrl}
                  alt={title}
                  width={140}
                  height={120}
                  className="w-[140px] h-[120px] object-cover rounded-lg"
                />
              </div> 
              <div className="flex-grow flex flex-col">
                <h2 className="text-xl font-bold mb-2">{title}</h2>
                <p className="text-sm text-gray-600">
                  {description}
                </p>
              </div>
            </div>
            <Button className="w-full bg-[#f87171] hover:bg-[#ef4444] text-white">
              よせがきを見る
            </Button>
          </CardContent>
        </Card>
        <div 
          className="absolute w-8 h-8 -top-3 left-10 
          before:content-[''] before:absolute before:w-8 before:h-8
          before:bg-white before:rotate-45 before:rounded-[3px]
          after:content-[''] after:absolute after:w-5 after:h-8
          after:bg-white after:rotate-45 after:rounded-[3px]"
        />
      </div>
    </div>
  )
}
