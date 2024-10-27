import React from 'react'
import { ChevronUp, ChevronDown } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Card } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Spot } from '@/api/@types'

type SpotListProps = {
  spots: Spot[]
  isOpen: boolean
  onToggle: () => void
}

const MessageSpot = ({ name }: { name: string }) => {
  return (
    <Card className="m-2 p-4 hover:bg-gray-50 transition-colors">
      <div className="flex items-center space-x-4">
        <div className="flex-grow">
          <h3 className="font-semibold text-lg">{name}</h3>
          <div className="flex items-center text-sm text-gray-500 space-x-4"></div>
        </div>
      </div>
    </Card>
  )
}

export function SpotList({ spots, isOpen, onToggle }: SpotListProps) {
  return (
    <Card
      className={`fixed left-0 right-0 bottom-0 bg-gradient-to-tl from-blue-400/40 via-blue-400 to-blue-500 backdrop-blur-lg rounded-t-xl shadow-lg transition-all duration-300 ease-in-out ${
        isOpen ? 'h-[50vh]' : 'max-h-[80px]'
      }`}
    >
      <div
        className="p-4 flex justify-between items-center cursor-pointer bg-gradient-to-l from-blue-400 to-blue-500 text-white rounded-t-xl"
        onClick={onToggle}
      >
        <h2 className="text-lg font-semibold">近くのスポット</h2>
        <Button variant="ghost" size="sm" className="text-white">
          {isOpen ? (
            <ChevronDown className="h-5 w-5" />
          ) : (
            <ChevronUp className="h-5 w-5" />
          )}
          <span className="sr-only">
            {isOpen ? 'リストを閉じる' : 'リストを開く'}
          </span>
        </Button>
      </div>
      {isOpen && (
        <ScrollArea className="h-[calc(100%-56px)] pb-10">
          {spots.map((spot) => (
            <MessageSpot key={spot.name} {...spot} />
          ))}
        </ScrollArea>
      )}
    </Card>
  )
}
