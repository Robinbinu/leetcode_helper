import { useState } from 'react';

interface CodeBlockProps {
  code: string;
  language: string;
}

export default function CodeBlock({ code, language }: CodeBlockProps) {
  const [copied, setCopied] = useState(false);

  const copyToClipboard = () => {
    navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="code-block relative">
      <pre className="bg-secondary-light p-4 rounded-md overflow-x-auto border border-gray-700">
        <code className="text-primary">{code}</code>
      </pre>
      <button
        onClick={copyToClipboard}
        className="copy-button absolute top-2 right-2 bg-primary text-secondary px-2 py-1 rounded text-xs font-medium hover:bg-primary-dark transition-opacity opacity-0 group-hover:opacity-100"
      >
        {copied ? 'Copied!' : 'Copy'}
      </button>
    </div>
  );
}
