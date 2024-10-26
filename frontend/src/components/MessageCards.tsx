'use client'

import { useState } from 'react'
import { Plus } from 'lucide-react'
import { motion } from 'framer-motion'
import { MdOutlineAddReaction } from 'react-icons/md'

type MessageCardsProps = {
  cards: Card[]
}

type Card = {
  id: number
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

const stamp_map = {
  thumbs_up: '👍',
  heart: '❤️',
}

export const MessageCards: React.FC<MessageCardsProps> = ({ cards }) => {
  const [cardList, setCardList] = useState(cards)

  const addCard = () => {
    const newCard = {
      id: cardList.length + 1,
      author_name: '匿名G',
      content: 'おなかすいた',
      color: `bg-${['blue', 'pink', 'orange', 'purple', 'yellow'][Math.floor(Math.random() * 5)]}-300`,
      stamps: [],
    }
    setCardList([...cardList, newCard])
  }

  return (
    <div className="p-4">
      {cardList.map((card, index) => (
        <motion.div
          key={card.id}
          className={`absolute w-64 px-3 py-2 rounded-lg shadow-lg ${card.color}`}
          initial={{ opacity: 0, y: 50 }}
          animate={{
            opacity: 1,
            y: index * 140,
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
          <h2 className="text-lg text-gray-900 leading-none">
            {card.author_name}
          </h2>
          <div className="w-full h-24 mt-3 bg-white bg-opacity-50 rounded p-2 text-gray-700 flex justify-center items-center">
            <div className="text-lg">{card.content}</div>
          </div>
          <div className="flex items-center mt-2 gap-2 h-6">
            <MdOutlineAddReaction />
            {card.stamps?.map((stamp) => (
              <div
                className={`rounded-full flex gap-1 px-1.5 py-0.5 text-sm ${stamp.is_reacted ? 'bg-blue-400 text-white' : 'bg-gray-100 text-gray-600'}`}
                key={stamp.type}
              >
                <div>{stamp_map[stamp.type as keyof typeof stamp_map]}</div>
                <div>{stamp.count}</div>
              </div>
            ))}
          </div>
        </motion.div>
      ))}
      <button
        onClick={addCard}
        className="z-50 fixed bottom-4 right-4 w-14 h-14 bg-blue-500 rounded-full flex items-center justify-center text-white shadow-lg hover:bg-blue-600 transition-colors"
      >
        <Plus size={24} />
      </button>
    </div>
  )
}
