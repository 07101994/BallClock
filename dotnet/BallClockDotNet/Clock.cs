using System;
using System.Collections.Generic;
using System.Linq;

namespace BallClockDotNet
{
	public class Clock
	{
		private Stack<int> ones = new Stack<int>(5);
		private Stack<int> fives = new Stack<int>(12);
		private Stack<int> hours = new Stack<int>(13);

		private Queue<int> queue = new Queue<int>();

		private int cycles = 0;
		private int size = 0;

		private int[] index;
		private int remaining;

		public Clock(int size)
		{
			this.size = size;

			hours.Push(-1);

			for (int i = 1; i <= size; i++)
				queue.Enqueue(i);
		}

		public int GetLCM()
		{
			var result = 1;

			// run until we know how often each ball returns to its origin
			index = new int[size];
			remaining = size;
			while (remaining > 0)
				Increment();
			
			// get the distinct lowest-common denominators
			var set = index.Distinct().ToArray();

			// calculate the lowest-common multiple for the set
			foreach (var val in set)
				result = LCM(result, val);

			return result;
		}

		/// <summary>
		/// Increments the ball clock by one minute
		/// </summary>
		private void Increment()
		{
			ones.Push(queue.Dequeue());

			if (ones.Count < 5) return;

			fives.Push(ones.Pop());
			while (ones.Count > 0)
				queue.Enqueue(ones.Pop());

			if (fives.Count < 12) return;

			var hour = fives.Pop();
			while (fives.Count > 0)
				queue.Enqueue(fives.Pop());

			if (hours.Count < 12)
			{
				hours.Push(hour);
			}
			else
			{
				// we've completed one more 12-hour cycle
				cycles++;

				while (hours.Count > 1)
					queue.Enqueue(hours.Pop());
			
				// 13th ball goes on last
				queue.Enqueue(hour);

				LogRepeats();
			}
		}

		/// <summary>
		/// Looks for balls that returned to their origin 
		/// and logs how many cycles it took
		/// </summary>
		private void LogRepeats()
		{
			int val = 0;

			// iterate through the queue and check the index of each ball
			for (int i = 0; i < size; i++)
			{
				val = queue.Dequeue();

				if (val == i + 1 && index[i] == 0)
				{
					// log when the ball returned to its origin
					remaining--;
					index[i] = cycles;
				}

				queue.Enqueue(val);
			}
		}

		/// <summary>
		/// Calculates the least-common multiple of two values
		/// </summary>
		private static int LCM(int a, int b)
		{
			int max = Math.Max(a, b);
			int min = Math.Min(a, b);

			for (int i = 1; i <= min; i++)
				if ((max * i) % min == 0)
					return i * max;

			return min;
		}
	}
}
