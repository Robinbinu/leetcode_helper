import { useState, useEffect } from 'react';
import CodeBlock from './CodeBlock';

interface SolutionDisplayProps {
  solution: {
    explanation: string;
    code: string;
    hints: string[];
    timeComplexity: string;
    spaceComplexity: string;
  } | null;
  error: string | null;
}

export default function SolutionDisplay({ solution, error }: SolutionDisplayProps) {
  const [activeTab, setActiveTab] = useState<'explanation' | 'code' | 'hints'>('explanation');
  const [language, setLanguage] = useState<string>('javascript');

  // Detect language from code if possible
  useEffect(() => {
    if (solution?.code) {
      if (solution.code.includes('function') || solution.code.includes('const')) {
        setLanguage('javascript');
      } else if (solution.code.includes('def ') || solution.code.includes('import ')) {
        setLanguage('python');
      } else if (solution.code.includes('public class') || solution.code.includes('public static')) {
        setLanguage('java');
      } else if (solution.code.includes('func ') && solution.code.includes('package ')) {
        setLanguage('go');
      }
    }
  }, [solution]);

  if (error) {
    return (
      <div className="bg-red-900 border border-red-700 rounded-lg p-6 mt-6 text-white">
        <h3 className="text-red-300 text-lg font-medium mb-2">Error</h3>
        <p className="text-red-100">{error}</p>
      </div>
    );
  }

  if (!solution) {
    return (
      <div className="bg-secondary text-white rounded-lg p-6 mt-6 border border-secondary-light">
        <div className="text-center py-8">
          <div className="text-primary text-5xl mb-4">{"{ }"}</div>
          <h3 className="text-xl font-medium">Enter a LeetCode problem to get started</h3>
          <p className="text-gray-400 mt-2">Your solution will appear here</p>
        </div>
      </div>
    );
  }

  return (
    <div className="bg-secondary text-white rounded-lg shadow-md mt-6">
      <div className="border-b border-gray-700">
        <nav className="flex -mb-px">
          <button
            onClick={() => setActiveTab('explanation')}
            className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
              activeTab === 'explanation'
                ? 'border-primary text-primary'
                : 'border-transparent text-gray-400 hover:text-gray-300 hover:border-gray-500'
            }`}
          >
            Explanation
          </button>
          <button
            onClick={() => setActiveTab('code')}
            className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
              activeTab === 'code'
                ? 'border-primary text-primary'
                : 'border-transparent text-gray-400 hover:text-gray-300 hover:border-gray-500'
            }`}
          >
            Solution Code
          </button>
          <button
            onClick={() => setActiveTab('hints')}
            className={`py-4 px-6 text-center border-b-2 font-medium text-sm ${
              activeTab === 'hints'
                ? 'border-primary text-primary'
                : 'border-transparent text-gray-400 hover:text-gray-300 hover:border-gray-500'
            }`}
          >
            Hints
          </button>
        </nav>
      </div>

      <div className="p-6">
        {activeTab === 'explanation' && (
          <div>
            <h3 className="text-lg font-medium mb-3 text-primary">Solution Explanation</h3>
            <div className="prose max-w-none text-white">
              <p className="whitespace-pre-line">{solution.explanation}</p>
            </div>
            
            <div className="mt-4 pt-4 border-t border-gray-700">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <h4 className="text-sm font-medium text-primary">Time Complexity</h4>
                  <p className="mt-1 text-sm text-gray-300">{solution.timeComplexity || "Not specified"}</p>
                </div>
                <div>
                  <h4 className="text-sm font-medium text-primary">Space Complexity</h4>
                  <p className="mt-1 text-sm text-gray-300">{solution.spaceComplexity || "Not specified"}</p>
                </div>
              </div>
            </div>
          </div>
        )}

        {activeTab === 'code' && (
          <div>
            <h3 className="text-lg font-medium mb-3 text-primary">Solution Code</h3>
            <CodeBlock code={solution.code} language={language} />
          </div>
        )}

        {activeTab === 'hints' && (
          <div>
            <h3 className="text-lg font-medium mb-3 text-primary">Helpful Hints</h3>
            {solution.hints && solution.hints.length > 0 ? (
              <ul className="list-disc pl-5 space-y-2">
                {solution.hints.map((hint, index) => (
                  <li key={index} className="text-gray-300">{hint}</li>
                ))}
              </ul>
            ) : (
              <p className="text-gray-400">No hints available.</p>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
