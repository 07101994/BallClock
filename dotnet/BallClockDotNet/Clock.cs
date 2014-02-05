using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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

		private Action<int[], Clock> onCycle;

		public Clock(int size, Action<int[], Clock> onCycle)
		{
			this.size = size;
			this.onCycle = onCycle;

			hours.Push(-1);

			for (int i = 1; i <= size; i++)
				queue.Enqueue(i);
		}

		public void Increment()
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
				while (hours.Count > 1)
					queue.Enqueue(hours.Pop());
			
				// 13th ball goes on last
				queue.Enqueue(hour);

				RaiseOnCycle();
			}
		}

		private void RaiseOnCycle()
		{
			cycles++;

			if (onCycle != null)
			{
				int[] order = new int[size];
				for (int i = 0; i < size; i++)
				{
					order[i] = queue.Dequeue();
					queue.Enqueue(order[i]);
				}

				onCycle.Invoke(order, this);
			}
		}

		public TimeSpan Time
		{
			get
			{
				int minutes = (fives.Count * 5) + ones.Count;
				return new TimeSpan(hours.Count, minutes, 0);
			}
		}

		public int Days
		{
			get
			{
				return cycles / 2;
			}
		}

		public int Cycles { get { return cycles; } }
	}
}
