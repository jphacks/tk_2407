'use client'

import { motion } from 'framer-motion'
import { MdOutlineAddReaction } from 'react-icons/md'
import { useRef } from 'react'

type MessageCardsProps = {
  cardList: Card[]
  setCardList: (cardList: Card[]) => void
}

export type Card = {
  id: string
  author_name: string
  content: string
  color: string
  stamps: Stamp[]
}

type Stamp = {
  type: string
  count: number
  is_reacted: boolean
}

const stampMap = {
  thumbs_up: '👍',
  heart: '❤️',
}

export const colors = [
  'bg-blue-300',
  'bg-pink-300',
  'bg-orange-300',
  'bg-purple-300',
  'bg-yellow-300',
]

export const MessageCards: React.FC<MessageCardsProps> = ({
  cardList,
  setCardList,
}) => {
  const reactionButtonRef = useRef<HTMLButtonElement>(null)

  const handleStampClick = (cardId: string, stampType: string) => {
    setCardList(
      cardList.map((card) =>
        card.id === cardId
          ? {
              ...card,
              stamps: card.stamps.map((stamp) =>
                stamp.type === stampType
                  ? stamp.is_reacted
                    ? { ...stamp, count: stamp.count - 1, is_reacted: false }
                    : { ...stamp, count: stamp.count + 1, is_reacted: true }
                  : stamp
              ),
            }
          : card
      )
    )
  }

  return (
    <>
      <div className="p-4">
        {cardList.map((card, index) => (
          <motion.div
            key={card.id}
            className={`absolute w-64 px-3 py-2 rounded-lg shadow-lg ${card.color}`}
            initial={{ opacity: 0, y: 50 }}
            animate={{
              opacity: 1,
              y: index * 120,
              x: index % 2 === 0 ? 20 : 100,
            }}
            transition={{
              type: 'spring',
              stiffness: 260,
              damping: 20,
              duration: 2,
            }}
            style={{
              zIndex: cardList.length - index,
            }}
          >
            <h2 className="text-md text-gray-900 leading-none">
              {card.author_name}
            </h2>
            <div className="w-full h-24 mt-2 bg-white bg-opacity-50 rounded p-2 text-gray-700 flex justify-center items-center">
              <div className="text-lg">{card.content}</div>
            </div>
            <div className="flex items-center mt-2 gap-2 h-4">
              <button ref={reactionButtonRef}>
                <MdOutlineAddReaction />
              </button>
              {card.stamps?.map((stamp) => (
                <button
                  className={`rounded-full flex gap-1 px-1.5 py-0.5 text-xs ${stamp.is_reacted ? 'bg-blue-400 text-white' : 'bg-gray-100 text-gray-600'}`}
                  key={stamp.type}
                  onClick={() => handleStampClick(card.id, stamp.type)}
                >
                  <div>{stampMap[stamp.type as keyof typeof stampMap]}</div>
                  <div>{stamp.count}</div>
                </button>
              ))}
            </div>
          </motion.div>
        ))}
      </div>
    </>
  )
}
