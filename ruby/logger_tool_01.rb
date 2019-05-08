class RunLogger
	def initialize(opts, logs = ['err','pre','log'])
		@msg = Array.new
		@opt = opts
		@logs = logs
	end

	def pre(msg)
		m = "Log:-NN-#{Time.now} ::#{msg}::"
		@msg.push m
		rec m, __method__.to_s
	end

	def log(msg)
		m = "Log:-NN-#{Time.now} ::#{msg}::"
		@msg.push m
		rec m, __method__.to_s
	end

	def err(msg)
		m = "Log:-EE-#{Time.now} ::#{msg}::"
		@msg.push m
		rec m, __method__.to_s
	end

	def wor(msg)
		m = "Log:-WW-#{Time.now} ::#{msg}::"
		@msg.push m
		rec m, __method__.to_s
	end

	def rec(m, w)
		@opt.each do |o|
			o.call(m) if @logs.include?(w)
		end
	end
end


#Log = RunLogger.new([lambda {|m| p m}, lambda {|m| File.open('./logger.log', 'a+') {|f| f.write(m + "\n")}}])

