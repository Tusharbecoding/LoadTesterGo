import type { Metadata } from "next";
import type { ReactNode } from "react";
import localFont from "next/font/local";
import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Distributed Load Tester Dashboard",
  description: "Monitor load testing metrics in real time",
};

interface LayoutProps {
  children: ReactNode;
}

export default function RootLayout({ children }: LayoutProps) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} bg-gray-100 text-gray-800 antialiased`}
      >
        <header className="bg-blue-600 text-white py-4 shadow">
          <div className="container mx-auto">
            <h1 className="text-3xl font-bold">LoadTesterGo</h1>
          </div>
        </header>
        <main className="flex-grow container mx-auto px-4 py-8">
          {children}
        </main>
        <footer className="bg-gray-800 absolute bottom-0 w-full text-white py-4">
          <div className="container mx-auto text-center">
            <p>&copy; 2024 LoadTesterGo - All rights reserved</p>
          </div>
        </footer>
      </body>
    </html>
  );
}
