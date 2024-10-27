import { useState } from 'react'
import { colors } from './MessageCards'

type MessageModalProps = {
  handleCreate: (content: string, color: string) => void
}

export const MessageModal: React.FC<MessageModalProps> = ({ handleCreate }) => {
  const [content, setContent] = useState<string>('')
  const [color] = useState<string>(colors[Math.floor(Math.random() * 5)])

  return (
    <>
      <div className="fixed inset-0 bg-slate-500 opacity-80 flex justify-center items-center z-40"></div>
      <div className="fixed inset-0 flex flex-col justify-center items-center gap-4 z-50">
        <span className="text-white font-semibold text-md">
          よせがきを作成中 . . .
        </span>
        <div className="p-4">
          <div className={`w-64 px-3 py-2 rounded-lg shadow-lg ${color}`}>
            <h2 className="text-md text-gray-900 leading-none">あなた</h2>
            <div className="w-full h-24 mt-2 bg-white bg-opacity-50 rounded p-2 text-gray-700 flex justify-center items-center">
              <textarea
                className="text-lg w-full h-full resize-none bg-transparent"
                value={content}
                onChange={(e) => setContent(e.target.value)}
              ></textarea>
            </div>
            <div className="flex justify-center items-center mt-2 gap-2 h-8">
              <button
                className="text-sm text-white bg-blue-400 py-1 px-3 rounded-md"
                onClick={() => handleCreate(content, color)}
              >
                作成
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
