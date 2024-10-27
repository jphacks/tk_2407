'use client'

import { Card, MessageCards } from '@/components/MessageCards'
import { MessageModal } from '@/components/MessageModal'
import { Plus } from 'lucide-react'
import { useState } from 'react'

const cards: Card[] = [
  {
    id: 'a',
    author_name: '匿名A',
    content: 'やっほー',
    color: 'bg-blue-300',
    stamps: [
      {
        type: 'thumbs_up',
        count: 4,
        is_reacted: true,
      },
      {
        type: 'heart',
        count: 2,
        is_reacted: false,
      },
    ],
  },
  {
    id: 'b',
    author_name: '匿名B',
    content: 'ようこそ！',
    color: 'bg-pink-300',
    stamps: [
      {
        type: 'thumbs_up',
        count: 1,
        is_reacted: true,
      },
    ],
  },
  {
    id: 'c',
    author_name: '匿名C',
    content: 'おなかすいた',
    color: 'bg-orange-300',
    stamps: [
      {
        type: 'heart',
        count: 1,
        is_reacted: false,
      },
    ],
  },
  {
    id: 'd',
    author_name: '匿名D',
    content: 'また来年くるぞ〜',
    color: 'bg-blue-300',
    stamps: [],
  },
]

export default function MessagePage() {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false)
  const [cardList, setCardList] = useState(cards)

  const addCard = ({ content, color }: { content: string; color: string }) => {
    const newCard = {
      id: Math.random().toString(36).substring(2) + Date.now().toString(36),
      author_name: 'あなた',
      content: content,
      color: color,
      stamps: [],
    }
    setCardList([...cardList, newCard])
  }

  const handleCreate = (content: string, color: string) => {
    setIsModalOpen(false)
    addCard({ content, color })
  }

  return (
    <>
      <div className="bg-yellow-200 h-screen">
        <MessageCards cardList={cardList} setCardList={setCardList} />
        <button
          className="z-30 fixed bottom-4 right-4 w-14 h-14 bg-blue-500 rounded-full flex items-center justify-center text-white shadow-lg hover:bg-blue-600 transition-colors"
          onClick={() => setIsModalOpen(true)}
        >
          <Plus size={24} />
        </button>
        {isModalOpen && <MessageModal handleCreate={handleCreate} />}
      </div>
    </>
  )
}
