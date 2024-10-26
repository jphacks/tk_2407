import React from 'react'

export function CustomPin({
  imageUrl,
  size = 100,
}: {
  imageUrl?: string
  size?: number
}) {
  const scale = size / 100
  const viewBoxHeight = 130
  const svgHeight = size * (viewBoxHeight / 100)

  return (
    <svg
      width={size}
      height={svgHeight}
      viewBox={`0 0 100 ${viewBoxHeight}`}
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <g transform={`scale(${scale})`}>
        <defs>
          <clipPath id="circleClip">
            <circle cx="50" cy="50" r="40" />
          </clipPath>
          <filter id="shadow" x="-20%" y="-20%" width="140%" height="140%">
            <feDropShadow dx="0" dy="3" stdDeviation="3" floodOpacity="0.3" />
          </filter>
        </defs>

        {/* ピンの影 */}
        <path
          d="M50 10C28.4 10 10 28.4 10 50C10 82.5 50 120 50 120C50 120 90 82.5 90 50C90 28.4 71.6 10 50 10Z"
          fill="#00000033"
          filter="url(#shadow)"
          transform="translate(0, 5)"
        />

        {/* ピンの本体 */}
        <path
          d="M50 0C22.4 0 0 22.4 0 50C0 87.5 50 130 50 130C50 130 100 87.5 100 50C100 22.4 77.6 0 50 0Z"
          fill="#FF8080"
        />

        {/* 白い円（画像の背景） */}
        <circle cx="50" cy="50" r="45" fill="white" />

        {/* 丸くクリップされた画像 */}
        <image
          href={imageUrl}
          x="5"
          y="5"
          width="90"
          height="90"
          clipPath="url(#circleClip)"
          opacity="0.7"
          preserveAspectRatio="xMidYMid slice"
        />

        {/* ピンの光沢効果 */}
        <ellipse
          cx="50"
          cy="30"
          rx="35"
          ry="20"
          fill="url(#gradient)"
          opacity="0.2"
        />

        <defs>
          <radialGradient
            id="gradient"
            cx="50%"
            cy="50%"
            r="50%"
            fx="50%"
            fy="50%"
          >
            <stop offset="0%" stopColor="white" stopOpacity="1" />
            <stop offset="100%" stopColor="#FF8080" stopOpacity="0" />
          </radialGradient>
        </defs>
      </g>
    </svg>
  )
}
